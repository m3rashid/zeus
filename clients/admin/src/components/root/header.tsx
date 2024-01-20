import { Component } from 'solid-js';

import Brand from './brand';
import ModeToggle from './modeToggle';
import GlobalSearch from './globalSearch';
import UserProfile from './userProfile';

const Header: Component = () => {
  return (
    <header class='flex h-10 w-full items-center justify-between px-2 border-b-2'>
      <Brand />
      <div class='flex items-center justify-center gap-2'>
        <GlobalSearch />
        <ModeToggle />
        <UserProfile />
      </div>
    </header>
  );
};

export default Header;
