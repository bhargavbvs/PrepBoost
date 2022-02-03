import React from 'react';
import { ProSidebar, Menu, MenuItem, SubMenu, SidebarFooter, SidebarHeader, SidebarContent } from 'react-pro-sidebar';
import {FaGem, FaHeart} from 'react-icons/fa'
import {RiContactsLine} from 'react-icons/gi'
import './SideBar.scss'


function SideBar() {
  return <div>
      <ProSidebar>
          <SidebarHeader>
              <h3>PREP BOOST</h3>
          </SidebarHeader>
          <SidebarContent>
          <Menu iconShape="square">
    <MenuItem icon={<FaGem />}>Topic Wise Questions</MenuItem>
    <MenuItem icon={<FaGem />}>Year Wise Questions</MenuItem>
    <MenuItem icon={<FaGem />}>Search</MenuItem>
    <MenuItem icon={<FaGem />}>Share With Friends</MenuItem>
    <MenuItem icon={<FaGem />}>Rate Our Website</MenuItem>
    <MenuItem icon={<FaGem />}>Logout</MenuItem>
    <MenuItem>Contact US</MenuItem>
    
  </Menu>
          </SidebarContent>
  
  <SidebarFooter><h3>Footer GIT HUB LINK</h3>


  </SidebarFooter>
</ProSidebar>;
  </div>;
}

export default SideBar;
