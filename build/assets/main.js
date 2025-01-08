/**
 * @var {Function} binding
 */
const binding = window.__BINDING__;



/**
 * MUST BE API OF GO-HYBRID
 * 
 * @param {string} name 
 * @returns {Function(...any[])}
 */
const bind = (object, method) => {
  return (...args) => {
    return binding(JSON.stringify({
      object: object,
      method: method,
      data: args
    }));
  }
}


const Math2 = {
  add: bind('Math2', 'Add'),
  multiple: bind('Math2', 'Multiple'),
}

document.addEventListener('click', e => {
  // battery('The title must change ' + Math.ceil(Math.random() * 1000000));

  console.log(Math2.add(1, 2))

});


// window.battery()
//       .then(level => console.log("System batter level is ", level))
//       .catch(() => {})