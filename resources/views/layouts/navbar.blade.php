<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="utf-8">
  <meta content="width=device-width, initial-scale=1.0" name="viewport">

  <title>Dev - @yield('title')</title>
  <meta content="" name="description">

  <meta content="" name="keywords">
  <meta name="csrf-token" content="">
  <!-- Favicons -->
  <link href="{{ asset('assets/img/favicon.png') }}" rel="icon">
  <link href="{{ asset('assets/img/apple-touch-icon.png') }}" rel="apple-touch-icon">

  <!-- Google Fonts -->
  <link href="https://fonts.googleapis.com/css?family=Open+Sans:300,300i,400,400i,600,600i,700,700i|Nunito:300,300i,400,400i,600,600i,700,700i|Poppins:300,300i,400,400i,500,500i,600,600i,700,700i" rel="stylesheet">

  <!-- Vendor CSS Files -->
  <link href="{{ asset('assets/vendor/bootstrap/css/bootstrap.min.css') }}" rel="stylesheet">
  <link href="{{ asset('assets/vendor/bootstrap-icons/bootstrap-icons.css') }}" rel="stylesheet">
  <link href="{{ asset('assets/vendor/aos/aos.css') }}" rel="stylesheet">
  <link href="{{ asset('assets/vendor/remixicon/remixicon.css') }}" rel="stylesheet">
  <link href="{{ asset('assets/vendor/swiper/swiper-bundle.min.css') }}" rel="stylesheet">
  <link href="{{ asset('assets/vendor/glightbox/css/glightbox.min.css') }}" rel="stylesheet">

  <link href="{{ asset('assets/vendor/others/css/bootstrap.min.css') }}" rel="stylesheet">
  <link href="{{ asset('assets/vendor/others/css/jquery.dataTables.min.css') }}" rel="stylesheet">
  <link href="{{ asset('assets/vendor/others/css/dataTables.bootstrap4.min.css') }}" rel="stylesheet">

  <!-- Template Main CSS File -->
  <link href="{{ asset('assets/css/style.css') }}" rel="stylesheet">

  <!-- =======================================================
  * Template Name: FlexStart - v1.4.0
  * Template URL: https://bootstrapmade.com/flexstart-bootstrap-startup-template/
  * Author: BootstrapMade.com
  * License: https://bootstrapmade.com/license/
  ======================================================== -->
</head>

<body>
@csrf

  <!-- ======= Header ======= -->
  <header id="header" class="header fixed-top">
    <div class="container-fluid container-xl d-flex align-items-center justify-content-between">

      <a href="{{ asset('') }}" class="logo d-flex align-items-center">
        <img src="{{ asset('/assets/img/logo.png') }}" alt="">
        <span>Dev</span>
      </a>

@guest
<nav id="navbar" class="navbar">
        <ul>
          <li><a href="{{ route('login') }}">Login</a></li>
          <li><a href="{{ route('register-user') }}">Register</a></li>
        </ul>
        <i class="bi bi-list mobile-nav-toggle"></i>
      </nav><!-- .navbar --> 

@else


      <nav id="navbar" class="navbar">
        <ul>
          <li><a class="nav-link scrollto active" href="{{ asset('') }}">Home</a></li>
          <li><a class="nav-link scrollto" href="{{ asset('about') }}">About</a></li>
          <li><a class="nav-link scrollto" href="{{ asset('') }}#services">Services</a></li>
          <li><a class="nav-link scrollto" href="{{ asset('') }}#portfolio">Portfolio</a></li>
          <li><a class="nav-link scrollto" href="{{ asset('') }}#team">Team</a></li>
          <li><a class="nav-link scrollto" href="{{ asset('blog') }}">Blog</a></li>
          <li class="dropdown"><a href="{{ asset('') }}#"><span>Dev</span> <i class="bi bi-chevron-down"></i></a>
            <ul>
              <li><a href="{{ asset('LaravelQuestions') }}">Laravel Questions</a></li>
              <li class="dropdown"><a href="#"><span>Laravel Dev</span> <i class="bi bi-chevron-right"></i></a>
                <ul>
                  <li><a href="{{ asset('LaravelDatabaseSeeding') }}">Laravel Database Seeding</a></li>
                  <li><a href="{{ asset('LaravelTinker') }}">Laravel Tinker</a></li>
                  <li><a href="#">Deep Drop Down 2</a></li>
                  <li><a href="#">Deep Drop Down 3</a></li>
                  <li><a href="#">Deep Drop Down 4</a></li>
                  <li><a href="#">Deep Drop Down 5</a></li>
                </ul>
              </li>
              <li><a href="#">Drop Down 2</a></li>
              <li><a href="#">Drop Down 3</a></li>
              <li><a href="#">Drop Down 4</a></li>
            </ul>
          </li>
          <li><a class="nav-link scrollto" href="{{ asset('company') }}">Company</a></li>
          <li class="dropdown"><a href="{{ asset('users') }}"><span>Users</span> <i class="bi bi-chevron-down"></i></a>
            <ul>
              <li><a href="{{ asset('usercreate') }}">Add new user</a></li>
              <li class="dropdown"><a href="#"><span>Deep Drop Down</span> <i class="bi bi-chevron-right"></i></a>
                <ul>
                  <li><a href="#">Drop Down 1</a></li>
                  <li><a href="#">Deep Drop Down 2</a></li>
                  <li><a href="#">Deep Drop Down 3</a></li>
                  <li><a href="#">Deep Drop Down 4</a></li>
                  <li><a href="#">Deep Drop Down 5</a></li>
                </ul>
              </li>
              <li><a href="#">Drop Down 2</a></li>
              <li><a href="#">Drop Down 3</a></li>
              <li><a href="#">Drop Down 4</a></li>
            </ul>
          </li>
          <li class="dropdown"><a href="#"><span>Profile</span> <i class="bi bi-chevron-down"></i></a>
            <ul>
@guest
                    <li><a href="{{ route('login') }}">Login</a>
                    </li>
                    <li><a href="{{ route('register-user') }}">Register</a>
                    </li>
@else
              <li><a href="{{ route('signout') }}">logout</a></li>
              <li class="dropdown"><a href="#"><span>Deep Drop Down</span> <i class="bi bi-chevron-right"></i></a>
                <ul>
                  <li><a href="#">Drop Down 1</a></li>
                  <li><a href="#">Deep Drop Down 2</a></li>
                  <li><a href="#">Deep Drop Down 3</a></li>
                  <li><a href="#">Deep Drop Down 4</a></li>
                  <li><a href="#">Deep Drop Down 5</a></li>
                </ul>
              </li>
          @endguest

            </ul>
          </li>
        </ul>
        <i class="bi bi-list mobile-nav-toggle"></i>
      </nav><!-- .navbar --> 
      
@endguest

    </div>




  </header>

  <!-- E n d Header -->