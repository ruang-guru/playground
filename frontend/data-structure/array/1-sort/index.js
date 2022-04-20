// Contoh Array: 23,98,67,45,12,24
// Prosedur: Ambil pasangan pertama dan tukar jika tidak berurutan dan ulangi sampai elemen terakhir diurutkan

// ###### iteration 1:
// step 1: 23,98,67,45,12,24
// step 2: 23,67,98,45,12,24
// step 3: 23,67,45,98,12,24
// step 4: 23,67,45,12,98,24
// step 5: 23,67,45,12,24,98

// result: we got last element sorted!!

// ###### iteration 2:
// step 1: 23,67,45,12,24,98
// step 2: 23,45,67,12,24,98
// step 3: 23,45,12,67,24,98
// step 4: 23,45,12,24,67,98

// result: we got second last element sorted!!

// ###### iteration 3:
// step 1: 23,45,12,24,67,98
// step 2: 23,12,45,24,67,98
// step 3: 23,12,24,45,67,98

// result: we got third last element sorted!!

// ###### iteration 4:
// step 1: 23,12,24,45,67,98
// step 2: 12,23,24,45,67,98

// hasil: kita mendapat array yang diurutkan!!

function sort(arr) {
    let swapped = false; //Untuk memeriksa apakah array sudah diurutkan; kemudian return;
    for (let i = 0; i < arr.length-1; i++) {
        for (let j = 0; j < arr.length-1; j++) {
            if (arr[j] > arr[j+1]) {
				//elemen bertukar
                let tmp = arr[j+1];
                arr[j+1] = arr[j];
                arr[j] = tmp;
                swapped = true;
            }
        }
        if (!swapped) {
            return arr;
        }
    }
    return arr;
}
let nums = [23, 98, 67, 45, 12, 24];
let sorted = sort(nums);
console.log(sorted);