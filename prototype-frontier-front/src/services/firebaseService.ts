import firebase from "firebase/app";
import "firebase/auth";

export default class FirebaseService {
  private static INSTANCE: FirebaseService;
  private user: firebase.User | null = null;
  constructor() {
    const firebaseConfig = {
      apiKey: process.env.VUE_APP_API_KEY,
      authDomain: process.env.VUE_APP_PROJECT_ID + ".firebaseapp.com",
      databaseURL:
        "https://" + process.env.VUE_APP_PROJECT_ID + ".firebaseio.com",
      projectId: process.env.VUE_APP_PROJECT_ID,
      storageBucket: process.env.VUE_APP_PROJECT_ID + ".appspot.com",
      messagingSenderId: process.env.VUE_APP_SENDER_ID,
      appId: process.env.VUE_APP_APP_ID,
      measurementId: "G-" + process.env.VUE_APP_MEASUREMENT_ID
    };
    if (firebase.apps.length === 0) {
      firebase.initializeApp(firebaseConfig);
    }
  }

  public static getInstance() {
    if (!FirebaseService.INSTANCE) {
      FirebaseService.INSTANCE = new FirebaseService();
    }
    return FirebaseService.INSTANCE;
  }

  public async getUser(forceReflesh = false): Promise<firebase.User | null> {
    if (!forceReflesh && this.user) {
      return this.user;
    }
    return new Promise<firebase.User | null>(resolve => {
      firebase.auth().onAuthStateChanged(user => {
        this.user = user;
        resolve(user);
      });
    });
  }

  public async getDisplayName() {
    return (await this.getUser())?.displayName;
  }

  public async getUID() {
    return (await this.getUser())?.uid;
  }

  public async getIdToken() {
    return await (await this.getUser(true))?.getIdToken();
  }

  public login() {
    const provider = new firebase.auth.GoogleAuthProvider();
    firebase.auth().signInWithRedirect(provider);
  }

  public async logout() {
    this.user = null;
    await firebase.auth().signOut();
  }
}
