// Importing Root and App using Promise.all (by using HTTP/2 multiplexing) loads them in parallel
declare global {
  namespace React {
    interface DOMAttributes<T> {
      onResize?: ReactEventHandler<T> | undefined;
      onResizeCapture?: ReactEventHandler<T> | undefined;
      nonce?: string | undefined;
    }
  }
}

Promise.all([import('@/Root'), import('@/App')]).then(([{ default: render }, { default: App }]) => {
  render(App);
});

export {};
