import { MessagePayload, Messaging, getMessaging, getToken, onMessage } from 'firebase/messaging';

import { app } from './firebase';

class NotificationService {
  messaging: Messaging;
  workerRegistration: ServiceWorkerRegistration | undefined;

  constructor() {
    this.messaging = getMessaging(app);
  }

  async registerWorker() {
    const registration = await navigator.serviceWorker.register('/firebase-messaging-sw.js', {
      type: 'module',
    });
    this.workerRegistration = registration;
  }

  async updateWorker() {
    this.workerRegistration?.update();
  }

  async getWebToken() {
    return await getToken(this.messaging, {
      vapidKey: import.meta.env.VITE_VAPID_KEY,
      serviceWorkerRegistration: this.workerRegistration,
    });
  }

  onNotifications(cb: (notification: MessagePayload) => void) {
    onMessage(this.messaging, (payload) => {
      console.info(`got notification`, payload);
      cb(payload);
    });
  }
}

export default new NotificationService();
