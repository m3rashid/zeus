import { Avatar } from '@admin/components/ui/avatar';
import { Button } from '@admin/components/ui/button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@admin/components/ui/dropdown-menu';
import { As } from '@kobalte/core';
import { Component } from 'solid-js';

const UserProfile: Component = () => {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <As component={Button} variant='ghost' size='sm' class='w-8 px-0'>
          <Avatar class='bg-red-300 w-8 h-8' />
          <span class='sr-only'>User Profile</span>
        </As>
      </DropdownMenuTrigger>
      <DropdownMenuContent>
        <DropdownMenuItem>
          <span>John Doe</span>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export default UserProfile;
