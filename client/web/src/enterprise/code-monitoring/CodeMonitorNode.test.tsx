import { mount } from 'enzyme'
import * as H from 'history'
import * as React from 'react'
import sinon from 'sinon'

import { CodeMonitorNode } from './CodeMonitoringNode'
import { mockCodeMonitor } from './testing/util'

describe('CreateCodeMonitorPage', () => {
    const history = H.createMemoryHistory()

    test('Does not show "Send test email" option when showCodeMonitoringTestEmailButton is false', () => {
        const component = mount(
            <CodeMonitorNode
                toggleCodeMonitorEnabled={sinon.spy()}
                location={history.location}
                node={mockCodeMonitor.node}
                isSiteAdminUser={true}
                showCodeMonitoringTestEmailButton={false}
            />
        )
        expect(component.find('.test-send-test-email').length).toBe(0)
    })

    test('Shows "Send test email" option to site admins on enabled code monitors', () => {
        const component = mount(
            <CodeMonitorNode
                toggleCodeMonitorEnabled={sinon.spy()}
                location={history.location}
                node={mockCodeMonitor.node}
                isSiteAdminUser={true}
                showCodeMonitoringTestEmailButton={true}
            />
        )
        expect(component.find('.test-send-test-email').length).toBe(1)
    })

    test('Does not show "Send test email" option when code monitor is disabled', () => {
        const component = mount(
            <CodeMonitorNode
                toggleCodeMonitorEnabled={sinon.spy()}
                location={history.location}
                node={{ ...mockCodeMonitor.node, enabled: false }}
                isSiteAdminUser={true}
                showCodeMonitoringTestEmailButton={true}
            />
        )
        expect(component.find('.test-send-test-email').length).toBe(0)
    })

    test('Does not show "Send test email" option to non-site admins', () => {
        const component = mount(
            <CodeMonitorNode
                toggleCodeMonitorEnabled={sinon.spy()}
                location={history.location}
                node={mockCodeMonitor.node}
                isSiteAdminUser={false}
                showCodeMonitoringTestEmailButton={true}
            />
        )
        expect(component.find('.test-send-test-email').length).toBe(0)
    })
})
