const user = {
    name:"No one",
    logincount:6, 
    signedin: true,

GetUserName: function() {
    console.log(`UserName is ${this.name}`);
    console.log(`UserName is ${this}`);
}
}
console.log(user.name)
console.log(user.GetUserName());