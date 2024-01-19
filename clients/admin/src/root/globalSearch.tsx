import { As } from '@kobalte/core';
import { TbSearch } from 'solid-icons/tb';
import { Component, createSignal } from 'solid-js';

import {
  Dialog,
  DialogTitle,
  DialogHeader,
  DialogContent,
  DialogTrigger,
  DialogDescription,
} from '@admin/components/ui/dialog';
import { Input } from '@admin/components/ui/input';
import { Button } from '@admin/components/ui/button';

const GlobalSearch: Component = () => {
  const [searchText, setSearchText] = createSignal('');

  const handleChange = (text: string) => {
    console.log(text);
    setSearchText(text);
  };

  return (
    <Dialog modal onOpenChange={() => setSearchText('')}>
      <DialogTrigger asChild>
        <As component={Button} variant='ghost' size='sm' class='w-9 px-0'>
          <TbSearch />
          <span class='sr-only'>Global Search</span>
        </As>
      </DialogTrigger>

      <DialogContent class='top-20'>
        <DialogHeader>
          <DialogTitle>
            <Input
              name='search'
              placeholder='Search ...'
              value={searchText()}
              onInput={(e) => handleChange(e.currentTarget.value)}
            />
          </DialogTitle>
          <DialogDescription>{/* Search Results */}</DialogDescription>
        </DialogHeader>
      </DialogContent>
    </Dialog>
  );
};

export default GlobalSearch;
