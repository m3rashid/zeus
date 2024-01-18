import { Component, For, Show } from 'solid-js';
import { SidebarRoute } from '../routes';

const Sidebar: Component<{ routes: SidebarRoute[] }> = (props) => {
  return (
    <ul class='space-y-2'>
      <For each={props.routes}>
        {(route) => (
          <>
            <Show when={route.children.length === 0}>
              <li>
                <a
                  href={route.path}
                  class='flex items-center p-2 text-base font-medium text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group'
                >
                  <route.icon />
                  <span class='ml-3'>{route.name}</span>
                </a>
              </li>
            </Show>

            <Show when={route.children.length > 0}>
              <li>
                <button
                  type='button'
                  class='flex items-center p-2 w-full text-base font-medium text-gray-900 rounded-lg transition duration-75 group hover:bg-gray-100 dark:text-white dark:hover:bg-gray-700'
                  aria-controls={'dropdown-' + route.name}
                  data-collapse-toggle={'dropdown-' + route.name}
                >
                  <route.icon />
                  <span class='flex-1 ml-3 text-left whitespace-nowrap'>
                    {route.name}
                  </span>

                  <svg
                    aria-hidden='true'
                    class='w-6 h-6'
                    fill='currentColor'
                    viewBox='0 0 20 20'
                    xmlns='http://www.w3.org/2000/svg'
                  >
                    <path
                      fill-rule='evenodd'
                      d='M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z'
                      clip-rule='evenodd'
                    ></path>
                  </svg>
                </button>

                <ul id={'dropdown-' + route.name} class='hidden py-2 space-y-2'>
                  <For each={route.children}>
                    {(child) => (
                      <li>
                        <a
                          href={child.path}
                          class='flex items-center p-2 pl-11 w-full text-base font-medium text-gray-900 rounded-lg transition duration-75 group hover:bg-gray-100 dark:text-white dark:hover:bg-gray-700'
                        >
                          {child.name}
                        </a>
                      </li>
                    )}
                  </For>
                </ul>
              </li>
            </Show>
          </>
        )}
      </For>
    </ul>
  );
};

export default Sidebar;
