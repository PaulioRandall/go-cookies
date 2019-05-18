# go-cookies
Collection of Go snippets, examples, and documentation.

The library was designed with the
[spike & stabalise](https://www.youtube.com/watch?v=lvs7VEsQzKY) development
approach in mind. During the spiking phase, the packages can be imported like
any other for quick development and testing. During the stabalisation phase, the
functions and files are copied in to the project source in order to remove the
dependency and allow the specifics of the project to change and refactor the
implementations as needed. Because the packages do not depend on one another
there shouldn't be any need to chase further dependencies.

The only exceptions are tests and the 'toastify' package. Tests can be copied as
well to aid refactoring and keep QA happy. The 'toastify' package is designed
for use with 'github.com/stretchr/testify' library by adding any missing
assertions and types I regularly use.