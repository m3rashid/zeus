import {
  Component,
  JSX,
  Show,
  children,
  mergeProps,
  splitProps,
} from 'solid-js';

export type ButtonProps = JSX.ButtonHTMLAttributes<HTMLButtonElement> & {
  beforeIcon?: JSX.Element;
  icon?: JSX.Element;
  buttonType?: 'primary' | 'secondary' | 'default';
  size?: 'xs' | 'sm' | 'md' | 'lg' | 'xl';
  rounded?: boolean;
  soft?: boolean;
  loading?: boolean;
};

const Button: Component<ButtonProps> = (props) => {
  const buttonChildren = children(() => props.children);
  const [local, others] = splitProps(props, [
    'icon',
    'size',
    'soft',
    'rounded',
    'beforeIcon',
    'buttonType',
  ]);

  const merged = mergeProps(
    {
      size: 'md',
      soft: false,
      rounded: false,
      loading: false,
      buttonType: 'default',
    },
    local
  );

  return (
    <button
      role='button'
      classList={{
        'flex items-center justify-between gap-2 font-semibold': true,
        'bg-blue-100 hover:bg-blue-200 text-blue-700':
          merged.buttonType === 'primary' && merged.soft,
        'bg-blue-500 hover:bg-blue-600 text-white':
          merged.buttonType === 'primary' && !merged.soft,

        'bg-gray-100 hover:bg-gray-200 text-gray-700':
          merged.buttonType === 'secondary' && merged.soft,
        'bg-gray-500 hover:bg-gray-600 text-white':
          merged.buttonType === 'secondary' && !merged.soft,

        'bg-gray-200 hover:bg-gray-300 text-gray-700':
          merged.buttonType === 'default',
        'px-2 py-1 text-xs': merged.size === 'xs',
        'px-3 py-1 text-sm': merged.size === 'sm',
        'px-4 py-2 text-sm': merged.size === 'md',
        'px-5 py-2 text-md': merged.size === 'lg',
        'px-6 py-3 text-md': merged.size === 'xl',
        'rounded-full': merged.rounded,
        'rounded-md': !merged.rounded,
        disabled: merged.loading,
      }}
      {...others}
    >
      <Show when={local.beforeIcon}>{local.beforeIcon}</Show>
      <div>{buttonChildren()}</div>
      <Show when={local.icon}>{local.icon}</Show>
    </button>
  );
};

export default Button;
