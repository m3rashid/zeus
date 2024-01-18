import { Component } from 'solid-js';

export const chartIcon = () => (
  <svg
    aria-hidden='true'
    class='w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white'
    fill='currentColor'
    viewBox='0 0 20 20'
    xmlns='http://www.w3.org/2000/svg'
  >
    <path d='M2 10a8 8 0 018-8v8h8a8 8 0 11-16 0z'></path>
    <path d='M12 2.252A8.014 8.014 0 0117.748 8H12V2.252z'></path>
  </svg>
);

export const pagesIcon = () => (
  <svg
    aria-hidden='true'
    class='flex-shrink-0 w-6 h-6 text-gray-500 transition duration-75 group-hover:text-gray-900 dark:text-gray-400 dark:group-hover:text-white'
    fill='currentColor'
    viewBox='0 0 20 20'
    xmlns='http://www.w3.org/2000/svg'
  >
    <path
      fill-rule='evenodd'
      d='M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2H6a2 2 0 01-2-2V4zm2 6a1 1 0 011-1h6a1 1 0 110 2H7a1 1 0 01-1-1zm1 3a1 1 0 100 2h6a1 1 0 100-2H7z'
      clip-rule='evenodd'
    ></path>
  </svg>
);

export type PlainRoute = {
  path: string;
  name: string;
};

export type SidebarRoute = PlainRoute & {
  icon: Component;
  children: PlainRoute[];
};

export const sidebarRoutesTop: SidebarRoute[] = [
  {
    path: '/',
    name: 'Overview',
    icon: chartIcon,
    children: [],
  },
  {
    path: '/home',
    name: 'Home',
    icon: chartIcon,
    children: [
      { name: 'Home1', path: '/' },
      { name: 'Home1', path: '/' },
    ],
  },
  {
    path: '/pages',
    name: 'Pages',
    icon: pagesIcon,
    children: [
      { name: 'Hello 1', path: '/world1' },
      { name: 'Hello 2', path: '/world2' },
      { name: 'Hello 3', path: '/world3' },
      { name: 'Hello 4', path: '/world4' },
    ],
  },
];

export const sidebarRoutesMiddle: SidebarRoute[] = [];
