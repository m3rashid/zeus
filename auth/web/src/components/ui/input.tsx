import type { Component, ComponentProps } from 'solid-js';
import { splitProps } from 'solid-js';

import { cn } from '@web/lib/utils';

const Input: Component<ComponentProps<'input'>> = (props) => {
  const [, rest] = splitProps(props, ['type', 'class']);
  return (
    <input
      type={props.type}
      class={cn(
        'border-input ring-offset-background placeholder:text-muted-foreground flex h-10 w-full rounded-md border bg-transparent px-3 py-2 text-sm file:border-0 file:bg-transparent file:text-sm file:font-medium disabled:cursor-not-allowed disabled:opacity-50',
        props.class
      )}
      {...rest}
    />
  );
};

export { Input };
