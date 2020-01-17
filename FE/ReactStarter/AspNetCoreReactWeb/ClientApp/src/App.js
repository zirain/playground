import React, { Component } from 'react';
import { Layout } from 'antd';
import { NavMenu } from './components/NavMenu';
const { Header, Content, Footer, Sider } = Layout;

export default class App extends Component {
  state = {
    collapsed: false,
  };

  onCollapse = (collapsed) => {
    console.log(collapsed);
    this.setState({ collapsed });
  }

  render() {
    return (
      <Layout style={{ minHeight: '100vh' }}>
        <Sider>
          <NavMenu />
        </Sider>
        <Layout>
          <Header style={{ background: '#fff', padding: 0 }} >
            <h1> React & Antd Demo</h1>
          </Header>
          <Content style={{ margin: '0 16px' }}>
            {this.props.children}
          </Content>
          <Footer style={{ textAlign: 'center' }}>
            Ant Design ©2018 Created by Ant UED
          </Footer>
        </Layout>
      </Layout>
    );
  }
}
