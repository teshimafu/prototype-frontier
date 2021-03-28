import firebase from "firebase/app";
import "firebase/auth";
import "firebase/database";
import firebaseConfig from "../../firebaseConfig.json";
import restAPIConfig from "../../restAPIConfig.json";

export default class FirebaseService {
  private static INSTANCE: FirebaseService;
  private user: firebase.User | null = null;
  constructor() {
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
    if (forceReflesh && this.user) {
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

  public async getAccessCount(): Promise<number | null> {
    const db = firebase.database();
    let count: number;
    try {
      const snapshot = await db
        .ref("settings/" + restAPIConfig.counterKey)
        .get();
      count = snapshot.val() + 1;
    } catch {
      return null;
    }
    db.ref("settings/" + restAPIConfig.counterKey).set(count);
    return count;
  }
}
