let palavra = "banana";
let resultado = "";

for (i = 0; i < palavra.length; i++) {
  if (palavra[i] === "a") {
    resultado += "o";
  } else {
    resultado += palavra[i];
  }
}
console.log(resultado);
