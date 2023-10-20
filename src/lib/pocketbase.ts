import PocketBase from "pocketbase";
import { writable } from "svelte/store";

export const pb = new PocketBase("https://pbserver.fly.dev/");

function userStore() {
    let unsubscribe: () => void;

    const { subscribe } = writable(pb.authStore.model ?? null, (set) => {
        unsubscribe = pb.authStore.onChange((auth) => {
            console.log("authStore changed", auth);
            set(pb.authStore.model);
        });
        return () => unsubscribe();
    });

    return {
        subscribe,
    };
}

export const user = userStore();
