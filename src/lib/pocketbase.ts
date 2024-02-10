import PocketBase from "pocketbase";
import { writable } from "svelte/store";
import { PB_ADDR } from "$env/static/private";
export const pb = new PocketBase(PB_ADDR);

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
