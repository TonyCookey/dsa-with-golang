const people = ['Nemo', 'Nema']
let t1 = performance.now()
for (let index = 0; index < people.length; index++) {
    if (people[index] == 'Nemo') {
        console.log('Found Nemo!')
    }

}
let t2 = performance.now()

console.log("Find Nemo took:", (t2 - t1), "milliseconds")