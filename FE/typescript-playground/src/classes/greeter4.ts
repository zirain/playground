class Student {
    fullName: string;
    constructor(public firstName, public middleInitial, public lastName) {
        this.fullName = firstName + " " + middleInitial + " " + lastName;
    }
}

interface IPerson {
    firstName: string;
    lastName: string;
}

function greeter4(person: IPerson) {
    return "Hello," + person.firstName + " " + person.lastName;
}

var user = new Student("Jane", "M.", "User");
document.body.innerHTML = greeter4(user);