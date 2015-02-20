var gulp = require("gulp");
var typescript = require("gulp-typescript");
var concat = require("gulp-concat");

gulp.task("tsc", function() {
    gulp.src(["./src/*.ts"])
        .pipe(typescript({target:"ES5", sortOutoput:true}))
        .js
        .pipe(concat("main.js"))
        .pipe(gulp.dest("./build/"));
});

gulp.task("default", ["tsc"]);
