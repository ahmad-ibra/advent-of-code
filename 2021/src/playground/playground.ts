export default function test() {
    let arr = ["abc", 123, 12, true]
    for (const x in arr) { // indexes of array
        console.log("for...in array", x)
    }
    for (const x of arr) { // values of array
        console.log("for...of array", x)
    }

    let obj = {
        "key1": "val1",
        2: 23,
        key3: false,
        "key4": {
            "sub1": 111,
            234: "sub val"
        }
    }
    for (const x in obj) { // keys of an object
        if (obj.hasOwnProperty(x)) {

        }
        console.log("for...in obj", x)
    }
    for (const x of Object.keys(obj)) { // keys of an object, might as well just use x in obj above, dont use it
        console.log("for...of Object.keys(obj)", x);
    }
    for (const x of Object.values(obj)) { // values of an object
        console.log("for...of Object.values(obj)", x);
    }
    for (const x in Object.keys(obj)) { // indexes in object (0, 1, 2, 3) since Ojbect.keys() returns an array, dont use it
        console.log("for...in Object.keys(obj)", x);
    }
    for (const x in Object.values(obj)) { // indexes in object (0, 1, 2, 3) since Ojbect.values() returns an array, dont use it
        console.log("for...in Object.values(obj)", x);
    }
    //for (const x of obj) { // invalid
    //    console.log("for...of obj", x)
    //}
}
