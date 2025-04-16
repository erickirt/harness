//  Copyright 2023 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package packages

import (
	"fmt"
	"net/http"

	middlewareauthn "github.com/harness/gitness/app/api/middleware/authn"
	"github.com/harness/gitness/registry/app/api/handler/generic"
	"github.com/harness/gitness/registry/app/api/handler/maven"
	"github.com/harness/gitness/registry/app/api/handler/npm"
	"github.com/harness/gitness/registry/app/api/handler/nuget"
	"github.com/harness/gitness/registry/app/api/handler/packages"
	"github.com/harness/gitness/registry/app/api/handler/python"
	"github.com/harness/gitness/registry/app/api/handler/rpm"
	"github.com/harness/gitness/registry/app/api/middleware"
	"github.com/harness/gitness/types/enum"

	"github.com/go-chi/chi/v5"
)

type Handler interface {
	http.Handler
}

/**
 * NewRouter creates a new router for the package API.
 * It sets up the routes and middleware for handling package-related requests.
 * Paths look like:
 * For all packages: /{rootIdentifier}/{registryName}/<package_type>/<package specific routes> .
 */
func NewRouter(
	packageHandler packages.Handler,
	mavenHandler *maven.Handler,
	genericHandler *generic.Handler,
	pythonHandler python.Handler,
	nugetHandler nuget.Handler,
	npmHandler npm.Handler,
	rpmHandler rpm.Handler,
) Handler {
	r := chi.NewRouter()

	r.Route("/{rootIdentifier}/{registryIdentifier}", func(r chi.Router) {
		r.Use(middleware.StoreOriginalURL)

		r.Route("/maven", func(r chi.Router) {
			r.Use(middleware.CheckMavenAuthHeader())
			r.Use(middlewareauthn.Attempt(packageHandler.GetAuthenticator()))
			r.Use(middleware.CheckMavenAuth())
			r.Use(middleware.TrackDownloadStatForMavenArtifact(mavenHandler))
			r.Use(middleware.TrackBandwidthStatForMavenArtifacts(mavenHandler))
			r.Get("/*", mavenHandler.GetArtifact)
			r.Head("/*", mavenHandler.HeadArtifact)
			r.Put("/*", mavenHandler.PutArtifact)
		})

		r.Route("/generic", func(r chi.Router) {
			r.Use(middlewareauthn.Attempt(packageHandler.GetAuthenticator()))
			r.Use(middleware.CheckAuth())
			r.Use(middleware.TrackDownloadStatForGenericArtifact(genericHandler))
			r.Use(middleware.TrackBandwidthStatForGenericArtifacts(genericHandler))

			r.Get("/*", genericHandler.PullArtifact)
			r.Put("/*", genericHandler.PushArtifact)
		})

		r.Route("/python", func(r chi.Router) {
			r.Use(middlewareauthn.Attempt(packageHandler.GetAuthenticator()))
			r.Use(middleware.CheckAuth())

			// TODO (Arvind): Move this to top layer with total abstraction
			r.With(middleware.StoreArtifactInfo(pythonHandler)).
				With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsUpload)).
				Post("/*", pythonHandler.UploadPackageFile)
			r.With(middleware.StoreArtifactInfo(pythonHandler)).
				With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsDownload)).
				Get("/files/{image}/{version}/{filename}", pythonHandler.DownloadPackageFile)
			r.With(middleware.StoreArtifactInfo(pythonHandler)).
				With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsDownload)).
				Get("/simple/{image}/", pythonHandler.PackageMetadata)
			r.Get("/simple/{image}", func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, r.URL.Path+"/", http.StatusMovedPermanently)
			})
		})

		r.Route("/{packageType}", func(r chi.Router) {
			r.Use(middlewareauthn.Attempt(packageHandler.GetAuthenticator()))
			r.Use(middleware.CheckAuth())
			r.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
				packageType := chi.URLParam(r, "packageType")
				http.Error(w, fmt.Sprintf("Package type '%s' is not supported", packageType), http.StatusNotFound)
			})
		})

		r.Route("/nuget", func(r chi.Router) {
			r.Use(middlewareauthn.Attempt(packageHandler.GetAuthenticator()))

			r.With(middleware.StoreArtifactInfo(nugetHandler)).
				With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsUpload)).
				Put("/", nugetHandler.UploadPackage)
			r.With(middleware.StoreArtifactInfo(nugetHandler)).
				With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsDownload)).
				Get("/package/{id}/{version}/{filename}", nugetHandler.DownloadPackage)
			r.With(middleware.StoreArtifactInfo(nugetHandler)).
				With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsDownload)).
				Get("/index.json", nugetHandler.GetServiceEndpoint)
		})

		r.Route("/npm", func(r chi.Router) {
			r.Use(middlewareauthn.Attempt(packageHandler.GetAuthenticator()))
			r.Use(middleware.CheckAuth())
			r.Route("/@{scope}/{id}", func(r chi.Router) {
				r.With(middleware.StoreArtifactInfo(npmHandler)).
					With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsUpload)).
					Put("/", npmHandler.UploadPackage)

				r.With(middleware.StoreArtifactInfo(npmHandler)).
					With(middleware.TrackDownloadStats(packageHandler)).
					With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsDownload)).
					Get("/-/{version}/@{scope}/{filename}", npmHandler.DownloadPackageFile)

				r.With(middleware.StoreArtifactInfo(npmHandler)).
					With(middleware.TrackDownloadStats(packageHandler)).
					With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsDownload)).
					Get("/-/@{scope}/{filename}", npmHandler.DownloadPackageFileByName)

				r.With(middleware.StoreArtifactInfo(npmHandler)).
					With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsDownload)).
					Head("/-/@{scope}/{filename}", npmHandler.HeadPackageFileByName)

				r.With(middleware.StoreArtifactInfo(npmHandler)).
					With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsDownload)).
					Get("/", npmHandler.GetPackageMetadata)

				r.With(middleware.StoreArtifactInfo(npmHandler)).
					With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsDelete)).
					Delete("/-/{version}/@{scope}/{filename}/-rev/{revision}", npmHandler.DeleteVersion)
			})

			r.Route("/{id}", func(r chi.Router) {
				r.With(middleware.StoreArtifactInfo(npmHandler)).
					With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsUpload)).
					Put("/", npmHandler.UploadPackage)

				r.With(middleware.StoreArtifactInfo(npmHandler)).
					With(middleware.TrackDownloadStats(packageHandler)).
					With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsDownload)).
					Get("/-/{version}/{filename}", npmHandler.DownloadPackageFile)

				r.With(middleware.StoreArtifactInfo(npmHandler)).
					With(middleware.TrackDownloadStats(packageHandler)).
					With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsDownload)).
					Get("/-/{filename}", npmHandler.DownloadPackageFileByName)

				r.With(middleware.StoreArtifactInfo(npmHandler)).
					With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsDownload)).
					Head("/-/{filename}", npmHandler.HeadPackageFileByName)
				r.With(middleware.StoreArtifactInfo(npmHandler)).
					With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsDownload)).
					Get("/", npmHandler.GetPackageMetadata)
				r.With(middleware.StoreArtifactInfo(npmHandler)).
					With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsDelete)).
					Delete("/-/{version}/{filename}/-rev/{revision}", npmHandler.DeleteVersion)
			})

			r.Route("/-/package/@{scope}/{id}/dist-tags", func(r chi.Router) {
				registerDistTagRoutes(r, npmHandler, packageHandler)
			})

			r.Route("/-/package/{id}/dist-tags", func(r chi.Router) {
				registerDistTagRoutes(r, npmHandler, packageHandler)
			})

			r.Route("/@{scope}/-rev/{revision}", func(r chi.Router) {
				registerRevisionRoutes(r, npmHandler, packageHandler)
			})

			r.Route("/{id}/-rev/{revision}", func(r chi.Router) {
				registerRevisionRoutes(r, npmHandler, packageHandler)
			})
		})
		r.Route("/rpm", func(r chi.Router) {
			r.Use(middlewareauthn.Attempt(packageHandler.GetAuthenticator()))
			r.Use(middleware.CheckAuth())
			r.With(middleware.StoreArtifactInfo(rpmHandler)).
				With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsUpload)).
				Put("/*", rpmHandler.UploadPackageFile)
			r.With(middleware.StoreArtifactInfo(rpmHandler)).
				With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsDownload)).
				Get("/repodata/{file}", rpmHandler.GetRepoData)
			r.With(middleware.StoreArtifactInfo(rpmHandler)).
				With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsDownload)).
				Get("/package/{name}/{version}/{architecture}/{file}", rpmHandler.DownloadPackageFile)
		})
	})

	return r
}

func registerDistTagRoutes(r chi.Router, npmHandler npm.Handler, packageHandler packages.Handler) {
	r.With(middleware.StoreArtifactInfo(npmHandler)).
		With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsUpload)).
		Get("/", npmHandler.ListPackageTag)

	r.With(middleware.StoreArtifactInfo(npmHandler)).
		With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsUpload)).
		Route("/{tag}", func(r chi.Router) {
			r.Put("/", npmHandler.AddPackageTag)
			r.Delete("/", npmHandler.DeletePackageTag)
		})
}

func registerRevisionRoutes(r chi.Router, npmHandler npm.Handler, packageHandler packages.Handler) {
	r.Use(middleware.StoreArtifactInfo(npmHandler))
	r.With(middleware.RequestPackageAccess(packageHandler, enum.PermissionArtifactsDelete)).
		Route("/", func(r chi.Router) {
			r.Delete("/", npmHandler.DeletePackage)
			r.Put("/", npmHandler.DeletePreview)
		})
}
