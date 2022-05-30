import create from "zustand"
import { persist, devtools } from "zustand/middleware"
import { produce } from "immer"

const useChatStore = create(persist(devtools((set) => ({
    messages: [],
    // whitout immer
    // addMessage: (message) => set((state) => ({ ...state, messages: [...state.messages, message] })),
    // with immer
    addMessage: (message) => set(produce((state) => {
        state.messages.push(message)
    })),
    clearMessages: () => set((state) => ({ ...state, messages: [] })),
})),{name: "chat"}));

export default useChatStore;