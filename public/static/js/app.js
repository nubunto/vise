(function() {
	var vise = {
		view: function() {
			return m('div', 'hello, world!');
		}
	};
	m.mount(document.querySelector('#app'), vise);
}());
