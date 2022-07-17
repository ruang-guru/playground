# Studi Kasus: API

Terdapat API yang dapat diakses untuk menambahkan data film, mendapatkan data film, dan semua data film. Diketahui data dari API sebagai berikut :

|                    | Tambah Data                   | Ambil Data                        |Ambil Semua Data                |
|---                 |---                            |---                                |---                             |
| Alamat             | <http://localhost:8090/movie> |<http://localhost:8090/movie/{id>} | <http://localhost:8090/movies> |
| Method             |POST                           | GET                               |GET                             |
| Request per detik  | 10                            | 25                                |20                              |
| Indikator performa | Waktu respon                  |Waktu respon                       |Waktu respon                    |

## Tugas

Lakukan load test untuk API tersebut.

1. Berdasarkan informasi yang diberikan, lakukan test ke API tersebut menggunakan vegeta, lalu simpan hasilnya dalam bentuk plot.

## Pertanyaan

Apakah informasi API yang diberikan sudah lengkap ? Apa alasannya.
