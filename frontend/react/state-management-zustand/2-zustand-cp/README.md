# Oversimplified E-commerce app with zustand
## Expected Outcome

https://user-images.githubusercontent.com/33412865/169922520-ee9a3305-5f3d-487c-be72-374144d982d5.mp4

## Challenge
- Gunakan middleware `immer` dan `devtools` untuk memudahkan pengerjaan
### Create item store

__file :__ `src/store/itemStore.js`


buat lah store untuk menyimpan item-item apa saja yang ada di halaman

key|type|description
---|----|-----------
items | array | untuk menyimpan item
addItem(item) | function | untuk menambahkan item baru
removeItem(itemId) | function | untuk menghapus item

__item object structure:__
```TS
item: {
    id: string,
    name: string,
    image: string,
    price: number,
    stock: number
}
```

### Create cart store
__file :__ `src/store/cartStore.js`

buat lah store untuk menyimpan item-item apa saja yang ada di cart
key|type|description
---|----|-----------
items | array | untuk menyimpan item
addItem(item) | function | untuk menambahkan item ke cart
removeItem(itemId) | function | untuk menghapus item dari cart
changeQuantity(itemId, quantity) | function | untuk mengubah jumlah item pada cart

__item object structure:__
```TS
item: {
    id: string,
    name: string,
    image: string,
    price: number,
    stock: number,
    quantity: number
}
```

### Custom Middleware
__file :__ `src/store/cartStore.js`

Buatlah custom middleware untuk melimit jumlah item pada cart agar tidak kurang dari 1 dan tidak melebihi dari stock yang ada

### Persistent Data
__file :__ `src/store/cartStore.js`, `src/store/itemStore.js`

Agar data tidak hilang saat melakukan `reload`, simpan state pada `localstorage`.

__localstorage keys:__
- `cart` : untuk menyimpan state `cart`
- `items` : untuk menyimpan state `items`

__Hint:__ gunakana middleware `persist`
