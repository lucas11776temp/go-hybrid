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
      data: args.map(arg => {
        if (typeof arg == "object") {
          try {
            return JSON.stringify(arg)
          } catch {
            return JSON.stringify({})
          }
        }
        return arg
      })
    }));
  }
}

// Math2 testing object
const Math2 = {
  addition: bind('Math2', 'Addition'),
}

// Movement testing object
const Movement = {
  change: bind('Movement', 'Change'),
}

// Click handler to pass event
document.addEventListener('click', e => {
  const p = Movement.change({
    longitude: Math.random() * 360,
    latitude: Math.random() * 180
  });

  console.log(p);

  p.then(v => console.log("VALUE: ", v));
});