import { generateHash } from "../random/gen-hash";

export function logoutSession() {
    window.location.href = "/api/v1/auth/logout?key=" + generateHash(32);
    return
}
