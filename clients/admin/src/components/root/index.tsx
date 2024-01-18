import { Component, For, JSX, Show } from 'solid-js';
import Brand from './brand';
import GlobalSearch from './globalSearch';
import Notifications from './notifications';
import Apps from './apps';
import UserMenu from './userProfile';
import { sidebarRoutesMiddle, sidebarRoutesTop } from '../routes';
import Sidebar from './sidebar';

const RootContainer: Component<JSX.HTMLAttributes<HTMLDivElement>> = (
  props
) => {
  return (
    <div class='antialiased bg-gray-50 dark:bg-gray-900'>
      <nav class='bg-white border-b border-gray-200 px-4 py-2.5 dark:bg-gray-800 dark:border-gray-700 fixed left-0 right-0 top-0 z-50'>
        <div class='flex flex-wrap justify-between items-center'>
          <div class='flex justify-start items-center'>
            <button
              data-drawer-target='drawer-navigation'
              data-drawer-toggle='drawer-navigation'
              aria-controls='drawer-navigation'
              class='p-2 mr-2 text-gray-600 rounded-lg cursor-pointer md:hidden hover:text-gray-900 hover:bg-gray-100 focus:bg-gray-100 dark:focus:bg-gray-700 focus:ring-2 focus:ring-gray-100 dark:focus:ring-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white'
            >
              <svg
                aria-hidden='true'
                class='w-6 h-6'
                fill='currentColor'
                viewBox='0 0 20 20'
                xmlns='http://www.w3.org/2000/svg'
              >
                <path
                  fill-rule='evenodd'
                  d='M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h6a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z'
                  clip-rule='evenodd'
                ></path>
              </svg>
              <svg
                aria-hidden='true'
                class='hidden w-6 h-6'
                fill='currentColor'
                viewBox='0 0 20 20'
                xmlns='http://www.w3.org/2000/svg'
              >
                <path
                  fill-rule='evenodd'
                  d='M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z'
                  clip-rule='evenodd'
                ></path>
              </svg>
              <span class='sr-only'>Toggle sidebar</span>
            </button>
            <Brand />
            <GlobalSearch position='header' />
          </div>
          <div class='flex items-center lg:order-2'>
            <button
              type='button'
              data-drawer-toggle='drawer-navigation'
              aria-controls='drawer-navigation'
              class='p-2 mr-1 text-gray-500 rounded-lg md:hidden hover:text-gray-900 hover:bg-gray-100 dark:text-gray-400 dark:hover:text-white dark:hover:bg-gray-700 focus:ring-4 focus:ring-gray-300 dark:focus:ring-gray-600'
            >
              <span class='sr-only'>Toggle search</span>
              <svg
                aria-hidden='true'
                class='w-6 h-6'
                fill='currentColor'
                viewBox='0 0 20 20'
                xmlns='http://www.w3.org/2000/svg'
                // aria-hidden='true'
              >
                <path
                  clip-rule='evenodd'
                  fill-rule='evenodd'
                  d='M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z'
                ></path>
              </svg>
            </button>

            <Notifications />
            <Apps />
            <UserMenu />
          </div>
        </div>
      </nav>

      {/* <!-- Sidebar --> */}
      <aside
        class='fixed top-0 left-0 z-40 w-64 h-screen pt-14 transition-transform -translate-x-full bg-white border-r border-gray-200 md:translate-x-0 dark:bg-gray-800 dark:border-gray-700'
        aria-label='Sidenav'
        id='drawer-navigation'
      >
        <div class='overflow-y-auto py-5 px-3 h-full bg-white dark:bg-gray-800'>
          <GlobalSearch position='sidebar' />
          <Sidebar routes={sidebarRoutesTop} />

          <Show when={sidebarRoutesMiddle.length > 0}>
            <div class='my-2 space-y-2 border-t border-gray-200 dark:border-gray-700' />
          </Show>

          <Sidebar routes={sidebarRoutesMiddle} />
        </div>

        {/* <div class='hidden absolute bottom-0 left-0 justify-center p-4 space-x-4 w-full lg:flex bg-white dark:bg-gray-800 z-20'>
          <a
            href='#'
            class='inline-flex justify-center p-2 text-gray-500 rounded cursor-pointer dark:text-gray-400 hover:text-gray-900 dark:hover:text-white hover:bg-gray-100 dark:hover:bg-gray-600'
          >
            <svg
              aria-hidden='true'
              class='w-6 h-6'
              fill='currentColor'
              viewBox='0 0 20 20'
              xmlns='http://www.w3.org/2000/svg'
            >
              <path d='M5 4a1 1 0 00-2 0v7.268a2 2 0 000 3.464V16a1 1 0 102 0v-1.268a2 2 0 000-3.464V4zM11 4a1 1 0 10-2 0v1.268a2 2 0 000 3.464V16a1 1 0 102 0V8.732a2 2 0 000-3.464V4zM16 3a1 1 0 011 1v7.268a2 2 0 010 3.464V16a1 1 0 11-2 0v-1.268a2 2 0 010-3.464V4a1 1 0 011-1z'></path>
            </svg>
          </a>

          <a
            href='#'
            data-tooltip-target='tooltip-settings'
            class='inline-flex justify-center p-2 text-gray-500 rounded cursor-pointer dark:text-gray-400 dark:hover:text-white hover:text-gray-900 hover:bg-gray-100 dark:hover:bg-gray-600'
          >
            <svg
              aria-hidden='true'
              class='w-6 h-6'
              fill='currentColor'
              viewBox='0 0 20 20'
              xmlns='http://www.w3.org/2000/svg'
            >
              <path
                fill-rule='evenodd'
                d='M11.49 3.17c-.38-1.56-2.6-1.56-2.98 0a1.532 1.532 0 01-2.286.948c-1.372-.836-2.942.734-2.106 2.106.54.886.061 2.042-.947 2.287-1.561.379-1.561 2.6 0 2.978a1.532 1.532 0 01.947 2.287c-.836 1.372.734 2.942 2.106 2.106a1.532 1.532 0 012.287.947c.379 1.561 2.6 1.561 2.978 0a1.533 1.533 0 012.287-.947c1.372.836 2.942-.734 2.106-2.106a1.533 1.533 0 01.947-2.287c1.561-.379 1.561-2.6 0-2.978a1.532 1.532 0 01-.947-2.287c.836-1.372-.734-2.942-2.106-2.106a1.532 1.532 0 01-2.287-.947zM10 13a3 3 0 100-6 3 3 0 000 6z'
                clip-rule='evenodd'
              ></path>
            </svg>
          </a>
          <div
            id='tooltip-settings'
            role='tooltip'
            class='inline-block absolute invisible z-10 py-2 px-3 text-sm font-medium text-white bg-gray-900 rounded-lg shadow-sm opacity-0 transition-opacity duration-300 tooltip'
          >
            Settings page
            <div class='tooltip-arrow' data-popper-arrow></div>
          </div>
        </div> */}
      </aside>

      <main class='p-4 md:ml-64 h-auto pt-20'>{props.children}</main>
    </div>
  );
};

export default RootContainer;
