[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]
![GitHub go.mod Go version][goversion-url]
![Go Report Card][goreport-url]
[![codecov](https://codecov.io/gh/michaelmcallister/sierpinski/branch/main/graph/badge.svg?token=MP9P5F9IRF)](https://codecov.io/gh/michaelmcallister/sierpinski)
<!-- PROJECT LOGO -->
<br />
<p align="center">
  <h3 align="center">Sierpiński triangle</h3>

  <img src="res/sierpinski.gif" />

  <p align="center">
      The Sierpiński triangle is a fractal attractive fixed set with the overall shape of an equilateral triangle, subdivided recursively into smaller equilateral triangles.
      <br />
      It is constructed by following these steps:
      <ol>
          <li>Take three points in a plane to form a triangle, you need not draw it.</li>
          <li>Randomly select any point inside the triangle and consider that your current position.
          </li>
          <li>Randomly select any one of the three vertex points.</li>
          <li>Move half the distance from your current position to the selected vertex.
          </li>
          <li>Plot the current position.</li>
          <li>Repeat from step 3.
          </li>
      </ol>
  </p>
</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[license-shield]: https://img.shields.io/github/license/michaelmcallister/sierpinski.svg?style=flat-square
[license-url]: https://github.com/michaelmcallister/sierpinski/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/mpmcallister
[goversion-url]: https://img.shields.io/github/go-mod/go-version/michaelmcallister/sierpinski
[goreport-url]: https://goreportcard.com/badge/gojp/goreportcard