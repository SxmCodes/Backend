# Interfaces.

1. Interfaces kya hoti hain?

Bhai, interface ek contract hota hai jo bolta hai:

"Agar tu mere rules follow karega toh tu mere gang ka part ban sakta hai, warna nikal!" 😎

Matlab agar ek struct kisi interface ke saare functions implement kar leta hai, toh wo us interface ka type ban jata hai.

Go me interfaces ka use hota hai polymorphism achieve karne ke liye, jisse ek function multiple types ke objects handle kar sake bina unka actual type jaane.

2. Interface Banane Ka Tareeka

Ek interface bas function signatures batata hai, par unka implementation struct ko khud likhna padta hai.

```
// Interface banaya jisme bola ki jo bhi shape hai usko area aur perimeter nikalna aana chahiye.
type shape interface {
    area() float64
    perimeter() float64
}
```
shape ek interface hai jo define karta hai ki koi bhi shape area() aur perimeter() function likhna padega.

Isme function ka signature hota hai, par implementation struct dega.

### Rectangular ka area function. 

```
type rect stuct {
width, heigh float64
}
```
yeh rect ek struct hai jo ek rectangle ko represent karta hai.
Iske paas width** aur **height naam ke do properties hain jo float64 type ki hain.

```
func (r rect) area() float64 {
    return r.width * r.height
}
```
and same way hm rectangle ka perimeter bhi nikal sakte hai, by puttin in perimeter in place of area. 

#### In terms of Circle 



