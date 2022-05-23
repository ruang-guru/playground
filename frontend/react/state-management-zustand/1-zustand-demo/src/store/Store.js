import create from 'zustand';

const useStore = create((set) => ({
    name: 'John Doe',
    address: 'Somewhere in the world',
    bio: 'I am a person who likes to do things and stuff and stuff and stuff',
    setName: (name) => set(() => ({ name })),
    setAddress: (address) => set(() => ({ address })),
    setBio: (bio) => set(() => ({ bio })),
}));

export default useStore;
export { useStore };