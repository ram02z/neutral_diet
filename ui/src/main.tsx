// Importing Root and App using Promise.all (by using HTTP/2 multiplexing) loads them in parallel
Promise.all([import('@/Root'), import('@/App')]).then(([{ default: render }, { default: App }]) => {
  render(App);
});

export {};
