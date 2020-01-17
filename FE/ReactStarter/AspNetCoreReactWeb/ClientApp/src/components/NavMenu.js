import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { Menu, Icon } from 'antd';

export class NavMenu extends Component {
  displayName = NavMenu.name

  render() {
    return (
      <Menu
        model="inline"
        theme="dark"
      >
        <Menu.Item>
          <Link to={'/'}>          
            <Icon type="home" theme="outlined" />Home
          </Link>
        </Menu.Item>
        <Menu.Item>
          <Link to={'/Counter'}>
            <Icon type="rise" theme="outlined" />Counter
          </Link>
        </Menu.Item>
        <Menu.Item>
          <Link to={'/fetchdata'}>
            <Icon type="table" theme="outlined" />Fetch Data
          </Link>
        </Menu.Item>
      </Menu>
    );
  }
}
