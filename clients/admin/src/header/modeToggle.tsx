import { As, useColorMode } from '@kobalte/core';
import { TbSun, TbMoon, TbDeviceLaptop } from 'solid-icons/tb';

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@admin/components/ui/dropdown-menu';
import { Button } from '@admin/components/ui/button';

const ModeToggle = () => {
  const { setColorMode } = useColorMode();

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <As component={Button} variant='ghost' size='sm' class='w-9 px-0'>
          <TbSun class='rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0' />
          <TbMoon class='absolute rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100' />
          <span class='sr-only'>Toggle theme</span>
        </As>
      </DropdownMenuTrigger>
      <DropdownMenuContent>
        <DropdownMenuItem onSelect={() => setColorMode('light')}>
          <TbSun class='mr-2 h-4 w-4' />
          <span>Light</span>
        </DropdownMenuItem>
        <DropdownMenuItem onSelect={() => setColorMode('dark')}>
          <TbMoon class='mr-2 h-4 w-4' />
          <span>Dark</span>
        </DropdownMenuItem>
        <DropdownMenuItem onSelect={() => setColorMode('system')}>
          <TbDeviceLaptop class='mr-2 h-4 w-4' />
          <span>System</span>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export default ModeToggle;
