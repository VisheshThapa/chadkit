import { pb } from "$lib/pocketbase";
import type { PageServerLoad } from "./$types";
import { error } from "@sveltejs/kit";
export const load = (async ({ params }) => {
    try {
        const confirm = await pb.collection("users").confirmVerification(
            params.TOKEN,
        );
        console.log(params.TOKEN);
        console.log(confirm);
        return {
            confirm: confirm,
        };
    } catch (e) {
        console.log(e);
        // redirect(301, '/login');
        throw error(401, "Unauthorized request!");
    }
}) satisfies PageServerLoad;
