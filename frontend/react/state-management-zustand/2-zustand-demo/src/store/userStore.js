import create from "zustand"
import { persist } from "zustand/middleware"
import { immer } from "zustand/middleware/immer"

const middlewareMinimumLength = config => (set, get, api) => config((args) => {
  set(args)
  if (get().leftUser.name.length < 5 || get().rightUser.name.length < 5) {
    throw new Error("Name must be at least 5 characters long")
  }
}, get, api)

const useUserStore = create(
  persist(
    middlewareMinimumLength(
      immer((set) => ({
        leftUser: {
          name: "Bebeb Kanna",
          image:
            "https://aramajapan.com/wp-content/uploads/2016/08/aramajapan_kanna-hashimoto.png",
        },
        rightUser: {
          name: "Ayang V",
          image:
            "https://awsimages.detik.net.id/visual/2020/12/30/kim-tae-hyung-v-bts.jpeg?w=650",
        },
        // whitout immer middleware
        // setLeftUser: (user) => set((state) => ({ ...state, leftUser: user })),
        // setRightUser: (user) => set((state) => ({ ...state, rightUser: user })),

        // with immer middleware
        setLeftUser: (user) =>
          set((state) => {
            state.leftUser = user
          }),
        setRightUser: (user) =>
          set((state) => {
            state.rightUser = user
          }),
      })),
    ),
    { name: "user" },
  ),
)

export default useUserStore
