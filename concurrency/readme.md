#### Problem 4
##### Concurrency Task Worker

Implementasi sebuah program untuk merangkum data museum di indonesia berdasarkan lokasi kabupaten/kota, simpan dalam file csv nama yang sesuai.

```
Kota Jakarta Pusat.csv
Kota Malang.csv
``` 

* Untuk sumber informasi gunakan API yang disediakan oleh **Open Data Indonesia** khusus data [museum](http://data.go.id/dataset/museum-indonesia)
* Gunakan *net/http* package untuk mengambil data dari API yang disediakan.
* Olah data mentah dari API jika diperlukan.
* Implementasi *concurrent* process untuk process mengambilan data dari API menggunakan *goroutine* . untuk ilustrasi [ilustrasi](https://talks.golang.org/2012/concurrency.slide)
* Batasi jumlah *concurrent* process, dengan mengaplikasikan *Queue* dan *Worker*.

ilustrasi:
```
grab -concurrent_limit=2 -output=/home/bayu/museum 
```

