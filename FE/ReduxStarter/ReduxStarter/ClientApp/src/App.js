import React from 'react';
import { Layout } from 'antd';
import { NavMenu } from './components/NavMenu';
const { Header, Content, Footer, Sider } = Layout;

export default props => (
  <Layout style={{ minHeight: '100vh' }}>
        <Sider>
          <NavMenu />
        </Sider>
        <Layout>
          <Header style={{ background: '#fff', padding: 0 }} >
            <h1> React & Antd Demo</h1>
          </Header>
          <Content style={{ margin: '0 16px' }}>
            {props.children}
          </Content>
          <Footer style={{ textAlign: 'center' }}>
            Ant Design Â©2018 Created by Ant UED
          </Footer>
        </Layout>
      </Layout>
);
