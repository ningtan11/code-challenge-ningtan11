import { useState } from "react";
import { Link } from "react-router-dom";
import { Sidebar, Menu, MenuItem, SubMenu } from "react-pro-sidebar";
import PeopleOutlinedIcon from "@mui/icons-material/PeopleOutlined";
import MenuOutlinedIcon from "@mui/icons-material/MenuOutlined";
import CurrencyExchangeIcon from '@mui/icons-material/CurrencyExchange';
import PriceCheckIcon from '@mui/icons-material/PriceCheck';
import TokenIcon from '@mui/icons-material/Token';
import AttachMoneyIcon from '@mui/icons-material/AttachMoney';
import AppsIcon from '@mui/icons-material/Apps';
import PollIcon from '@mui/icons-material/Poll';
import WavesIcon from '@mui/icons-material/Waves';
import AnalyticsIcon from '@mui/icons-material/Analytics';
import LinkIcon from '@mui/icons-material/Link';
import SwapVerticalCircleIcon from '@mui/icons-material/SwapVerticalCircle';
import MasksIcon from '@mui/icons-material/Masks';
import WalletIcon from '@mui/icons-material/Wallet';
import MonetizationOnIcon from '@mui/icons-material/MonetizationOn';
import LiveHelpIcon from '@mui/icons-material/LiveHelp';
import "./SideNavbar.css";

function SideNavbar() {
  const [ collapsed, setCollapsed ] = useState(false);

  const collapseSidebar = () => {
    setCollapsed(!collapsed);
  }

  return (
    <div style={({ display: "flex" })}>
      <Sidebar collapsed={collapsed}>
        <Menu>
          <MenuItem className="menuitem"
            icon={<MenuOutlinedIcon />}
            onClick={collapseSidebar}
            style={{ textAlign: "center" }}
          >
            <h2>Fancy Form</h2>
          </MenuItem>
          <MenuItem className="menuitem" icon={<CurrencyExchangeIcon />} component={<Link to="/" />}>Swap</MenuItem>
          <SubMenu className="menuitem" icon={<PriceCheckIcon />} label="Prices">
            <MenuItem className="menuitem" icon={<AttachMoneyIcon />} component={<Link to="/prices/token" />}>Token</MenuItem>
            <MenuItem className="menuitem" icon={<TokenIcon />} component={<Link to="/prices/nfts" />}>NFTs</MenuItem>
          </SubMenu>
          <SubMenu className="menuitem" icon={<AppsIcon />} label="App">
            <MenuItem className="menuitem" icon={<PollIcon />} component={<Link to="/app/vote" />}>Vote</MenuItem>
            <MenuItem className="menuitem" icon={<WavesIcon />} component={<Link to="/app/pools" />}>Pools</MenuItem>
            <MenuItem className="menuitem" icon={<AnalyticsIcon />} component={<Link to="/app/analytics" />}>Analytics</MenuItem>
          </SubMenu>
          <SubMenu className="menuitem" icon={<LinkIcon />} label="Connect">
            <MenuItem className="menuitem" icon={<SwapVerticalCircleIcon/>} component={<Link to="/connect/uniswap" />}>Uniswap Wallet</MenuItem>
            <MenuItem className="menuitem" icon={<MasksIcon />} component={<Link to="/connect/metamask" />}>Install MetaMask</MenuItem>
            <MenuItem className="menuitem" icon={<WalletIcon />} component={<Link to="/connect/walletconnect" />}>WalletConnect</MenuItem>
            <MenuItem className="menuitem" icon={<MonetizationOnIcon />} component={<Link to="/connect/coinbase" />}>Coinbase Wallet</MenuItem>
          </SubMenu>
          <MenuItem className="menuitem" icon={<PeopleOutlinedIcon />} component={<Link to="/aboutUs" />}>About Us</MenuItem>
          <MenuItem className="menuitem" icon={<LiveHelpIcon />} component={<Link to="/faq" />}>FAQ</MenuItem>
        </Menu>
      </Sidebar>
    </div>
  );
}

export default SideNavbar;