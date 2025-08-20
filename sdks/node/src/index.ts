import { USM } from "./usm";

export async function load() {
  return await USM.load();
}

export { USM };