import React, { useRef } from 'react'
import { Container, HarnessDocTooltip, Layout, Text, TextInput } from '@harnessio/uicore'
import { Color } from '@harnessio/design-system'
import type { Cell, CellValue, ColumnInstance, Renderer, Row, TableInstance } from 'react-table'
import { Icon } from '@harnessio/icons'
import { cloneDeep } from 'lodash-es'
import { learnMoreRegion, type regionProp } from 'cde-gitness/constants'
import { useStrings } from 'framework/strings'
import RegionTable from 'cde-gitness/components/RegionTable/RegionTable'
import css from './InfraDetails.module.scss'

type CellTypeWithActions<D extends Record<string, any>, V = any> = TableInstance<D> & {
  column: ColumnInstance<D>
  row: Row<D>
  cell: Cell<D, V>
  value: CellValue<V>
}

type CellType = Renderer<CellTypeWithActions<any>>

interface customCellProps {
  column: {
    id: string
    placeholder: string
  }
  row: {
    index: number
  }
  value: string
}

interface LocationProps {
  regionData: regionProp[]
  setRegionData: (result: regionProp[]) => void
  initialData: regionProp
}

const ConfigureLocations = ({ regionData, setRegionData, initialData }: LocationProps) => {
  const { getString } = useStrings()
  const lastFocusRef = useRef<HTMLInputElement | null>(null)
  const currentFocusRef = useRef('')

  const deleteRegion = (indx: number) => {
    const clonedData = cloneDeep(regionData)
    const result: regionProp[] = []
    clonedData.forEach((region: regionProp, index: number) => {
      if (index !== indx) {
        result.push(region)
      }
    })
    setRegionData(result)
  }

  const ActionCell: CellType = (row: any) => {
    return (
      <Container className={css.deleteContainer}>
        <Icon name="code-delete" size={24} onClick={() => deleteRegion(row?.row?.index)} />
      </Container>
    )
  }

  const inputHandler = (key: string, value: Unknown, index: number) => {
    const clonedData = cloneDeep(regionData)
    const result: regionProp[] = clonedData?.map((region: regionProp, indx: number) => {
      if (index === indx) {
        region = {
          ...region,
          [key]: value
        }
      }
      return region
    })
    setRegionData(result)

    // Wait for the next render cycle to set focus
    setTimeout(() => {
      const parentNode: any = lastFocusRef?.current?.childNodes?.[0]
      const inputNode = parentNode?.querySelector('input')
      inputNode?.focus()
    }, 0)
  }

  const CustomCell: any = (row: customCellProps) => {
    const { id, placeholder } = row?.column
    const focusId = `${id}_${row?.row?.index}`
    return (
      <Container className={css.inputContainer} ref={currentFocusRef?.current === focusId ? lastFocusRef : null}>
        <TextInput
          placeholder={placeholder}
          name={id}
          value={row?.value}
          onFocus={() => {
            currentFocusRef.current = focusId ?? ''
          }}
          onChange={(e: any) => inputHandler(id, e.target.value, row?.row?.index)}
        />
      </Container>
    )
  }
  const columns = [
    {
      Header: (
        <Layout.Horizontal>
          <Text className={css.headingText}>{getString('cde.gitspaceInfraHome.region')}</Text>
          <HarnessDocTooltip tooltipId="InfraProviderRegionLocation" useStandAlone={true} />
        </Layout.Horizontal>
      ),
      Cell: CustomCell,
      accessor: 'location',
      placeholder: 'e.g us-west1',
      width: '13%'
    },
    {
      Header: (
        <Layout.Horizontal>
          <Text className={css.headingText}>{getString('cde.gitspaceInfraHome.defaultSubnet')}</Text>
          <HarnessDocTooltip tooltipId="InfraProviderRegionDefaultSubnet" useStandAlone={true} />
        </Layout.Horizontal>
      ),
      Cell: CustomCell,
      accessor: 'defaultSubnet',
      placeholder: 'e.g 10.6.0.0/16',
      width: '22%'
    },
    {
      Header: (
        <Layout.Horizontal>
          <Text className={css.headingText}>{getString('cde.gitspaceInfraHome.proxySubnet')}</Text>
          <HarnessDocTooltip tooltipId="InfraProviderRegionProxySubnet" useStandAlone={true} />
        </Layout.Horizontal>
      ),
      Cell: CustomCell,
      accessor: 'proxySubnet',
      placeholder: 'e.g 10.3.0.0/16',
      width: '22%'
    },
    {
      Header: (
        <Layout.Horizontal>
          <Text className={css.headingText}>{getString('cde.configureInfra.domain')}</Text>
          <HarnessDocTooltip tooltipId="InfraProviderRegionDomain" useStandAlone={true} />
        </Layout.Horizontal>
      ),
      Cell: CustomCell,
      accessor: 'domain',
      placeholder: 'e.g us-west-ga.io',
      width: '15%'
    },
    {
      Header: (
        <Layout.Horizontal>
          <Text className={css.headingText}>{getString('cde.gitspaceInfraHome.dnsManagedZone')}</Text>
          <HarnessDocTooltip tooltipId="InfraProviderRegionDNS" useStandAlone={true} />
        </Layout.Horizontal>
      ),
      Cell: CustomCell,
      accessor: 'dns',
      placeholder: 'e.g us-west-ga.io',
      width: '20%'
    },
    {
      Header: '',
      accessor: 'identifier',
      Cell: ActionCell,
      width: '8%'
    }
  ]

  const addNewRegion = () => {
    const clonedData: regionProp[] = cloneDeep(regionData)
    const payload: regionProp = {
      ...initialData,
      identifier: clonedData?.length + 1
    }
    clonedData.push(payload)
    setRegionData(clonedData)
  }

  return (
    <Layout.Vertical spacing="none" className={css.containerSpacing}>
      <Text className={css.basicDetailsHeading}>{getString('cde.configureInfra.configureLocations')}</Text>
      <Layout.Horizontal spacing="small" className={css.bottomSpacing}>
        <Text color={Color.GREY_400} className={css.headerLinkText}>
          {getString('cde.configureInfra.configureLocationNote')}
        </Text>
        <Text
          color={Color.PRIMARY_7}
          className={css.headerLinkText}
          onClick={() => {
            window.open(learnMoreRegion, '_blank')
          }}>
          {getString('cde.configureInfra.learnMore')}
        </Text>
      </Layout.Horizontal>

      <RegionTable columns={columns} addNewRegion={addNewRegion} regionData={regionData} />
    </Layout.Vertical>
  )
}

export default ConfigureLocations
