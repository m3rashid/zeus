/** @type {import('tailwindcss').Config} */
export default {
	darkMode: 'class',
	content: ["./**/*.{html,js,jsx,ts,tsx}", "./node_modules/flowbite/**/*.js"],
	theme: {
		extend: {},
	},
	plugins: [
		require('flowbite/plugin'),
	],
}

