module.exports = Util;

function Util() { }

Util.toUpper = function (s) {
    return s.substring(0, 1).toUpperCase() + s.substring(1);
};
