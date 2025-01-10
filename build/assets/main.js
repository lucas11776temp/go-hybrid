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
  addition: bind('Math2', 'Addition'),
  multiple: bind('Math2', 'Multiple'),
}

document.addEventListener('click', e => {
  const p = Math2.addition(1, 2);

  console.log(p);

  p.then(v => console.log("VALUE: ", v));
});