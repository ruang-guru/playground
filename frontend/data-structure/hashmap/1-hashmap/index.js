function printEmployee(employees, id) {
    if (employees.has(id)) {
        console.log(employees.get(id));
    } else {
        console.log('No employee with id ' + id);
    }
}
function isEmployeeExists(employees, id) {
    return employees.has(id);
}

let employees = new Map();
employees.set(1001, "John");
employees.set(1002, "Steve");
employees.set(1003, "Maria");

printEmployee(employees, 1001);
printEmployee(employees, 1010);
if (isEmployeeExists(employees, 1002)) {
    console.log("Employee with id 1002 exists");
}