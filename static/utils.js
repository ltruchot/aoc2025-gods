// Shared utilities for AoC 2025

// Promisified timeout
export const delay = ms => new Promise(resolve => setTimeout(resolve, ms));

// getElementById shorthand
export const byId = id => document.getElementById(id);

// Pick random element from array
export const pickRandom = arr => arr[Math.floor(Math.random() * arr.length)];

// Create element with attributes (curried)
export const createEl = R.curry((tag, attrs) => {
	const el = document.createElement(tag);
	Object.assign(el, attrs);
	return el;
});

export const div = createEl('div');
