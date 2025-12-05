// Shared utilities for AoC 2025
const U = {
	// Promisified timeout
	delay: ms => new Promise(resolve => setTimeout(resolve, ms)),

	// getElementById shorthand
	byId: id => document.getElementById(id),

	// Pick random element from array
	pickRandom: arr => arr[Math.floor(Math.random() * arr.length)],

	// Create element with attributes (curried)
	createEl: R.curry((tag, attrs) => {
		const el = document.createElement(tag);
		Object.assign(el, attrs);
		return el;
	}),

	// Shorthand for div
	div: null
};
U.div = U.createEl('div');
