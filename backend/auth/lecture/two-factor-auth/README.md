#Two Factor Authentication

##Requirement
1. Google Authenticator or other MFA [[here](https://play.google.com/store/apps/details?id=com.google.android.apps.authenticator2&hl=in&gl=US)]

##Step
1. Register MFA manual di Aplikasi Google Authenticator menggunakan secret yang sama dengan di `function VerifyOTP`\
contoh : `2MXGP5X3FVUEK6W4UB2PPODSP2GKYWUT`
2. Login ke dalam service dengan endpoint `/signin` (Tidak perlu akun karna sudah dimock)
3. Gunakan return dari endpoint sebelumnya sebagai header berikut\ 
`Authorization : Bearer <token>`'
4. Buka aplikasi MFA, kemudian masukan token ke endpoint `/verify-otp`