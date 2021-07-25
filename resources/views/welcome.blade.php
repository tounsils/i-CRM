@extends('layouts.main')

@section('title', 'Home')

@section('content')

    <!-- start reusable part -->

  <!-- ======= Hero Section ======= -->
  <section id="hero" class="hero d-flex align-items-center">

    <div class="container">
      <div class="row">
        <div class="col-lg-6 d-flex flex-column justify-content-center">
          <h1 data-aos="fade-up">We offer modern solutions for growing your business</h1>
          <h2 data-aos="fade-up" data-aos-delay="400">We are team of talented designers 
            making websites and web applications.</h2>
          <div data-aos="fade-up" data-aos-delay="600">
            <div class="text-center text-lg-start">
              <a href="mailto:tn76contact@gmail.com?subject=Contacting throu tn76.com&body=I need you to ..." class="btn-get-started scrollto d-inline-flex align-items-center justify-content-center align-self-center">
                <span>Get Started</span>
                <i class="bi bi-arrow-right"></i>
              </a>
            </div>
          </div>
        </div>
        <div class="col-lg-6 hero-img" data-aos="zoom-out" data-aos-delay="200">
          <img src="{{ asset('/assets/img/hero-img.png') }}" class="img-fluid" alt="">
        </div>
      </div>
    </div>

  </section><!-- End Hero -->

  <main id="main">
    <!-- ======= About Section ======= -->
    <section id="about" class="about">

      <div class="container" data-aos="fade-up">
        <div class="row gx-0">

          <div class="col-lg-6 d-flex flex-column justify-content-center" data-aos="fade-up" data-aos-delay="200">
            <div class="content">
              <h3>Who We Are</h3>
              <h2> The best of web design and web design inspiration.</h2>
              <p>
                Providing new designs and web development, and offering the best web solutions for 
                small to medium business. For the design and development of any tailored website 
                contact us.<br>
                The best of web design and web design inspiration 
                - Providing new designs and web development, and offering the best web 
                solutions for small to medium business. <br>
                For the design and development of any tailored website contact us.              </p>
              <div class="text-center text-lg-start">
                <a href="mailto:tn76contact@gmail.com?subject=Contacting throu tn76.com&body=I need you to ..." class="btn-read-more d-inline-flex align-items-center justify-content-center align-self-center">
                  <span>Contact us</span>
                  <i class="bi bi-arrow-right"></i>
                </a>
              </div>
            </div>
          </div>

          <div class="col-lg-6 d-flex align-items-center" data-aos="zoom-out" data-aos-delay="200">
            <img src="{{ asset('/assets/img/about.jpg') }}" class="img-fluid" alt="">
          </div>

        </div>
      </div>

    </section><!-- End About Section -->

    <!-- ======= Values Section ======= -->
    <section id="values" class="values" name="Our Web Design Company’s Values">

      <div class="container" data-aos="fade-up">

        <header class="section-header">
          <h2>Our Values</h2>
          <p>Our Web Design Company’s Values </p>
        </header>

        <div class="row">

          <div class="col-lg-4">
            <div class="box" data-aos="fade-up" data-aos-delay="200">
              <img src="{{ asset('/assets/img/values-1.png') }}" class="img-fluid" alt="">
              <h3>Friendly</h3>
              <p>Our approach is down-to-earth and just plain friendly. We love our clients, we love seeing you succeed, and we’re here to help..</p>
            </div>
          </div>

          <div class="col-lg-4 mt-4 mt-lg-0">
            <div class="box" data-aos="fade-up" data-aos-delay="400">
              <img src="{{ asset('/assets/img/values-2.png') }}" class="img-fluid" alt="">
              <h3>Thoughtful</h3>
              <p>We take an inquisitive, methodical approach to every project to ensure that we find the best solution possible. 
                Keeping our clients’ best interests at heart is a way of life for us..</p>
            </div>
          </div>

          <div class="col-lg-4 mt-4 mt-lg-0">
            <div class="box" data-aos="fade-up" data-aos-delay="600">
              <img src="{{ asset('/assets/img/values-3.png') }}" class="img-fluid" alt="">
              <h3>Intuitive</h3>
              <p>Technology can be frustratingly complex. We aim to keep things as painless as possible. 
                Our goal is that your experience will be simple and enjoyable..</p>
            </div>
          </div>

        </div>

      </div>

    </section><!-- End Values Section -->

    <!-- ======= Counts Section ======= -->
    <section id="counts" class="counts">
      <div class="container" data-aos="fade-up">

        <div class="row gy-4">

          <div class="col-lg-3 col-md-6">
            <div class="count-box">
              <i class="bi bi-emoji-smile"></i>
              <div>
                <span data-purecounter-start="0" data-purecounter-end="232" data-purecounter-duration="1" class="purecounter"></span>
                <p>Happy Clients</p>
              </div>
            </div>
          </div>

          <div class="col-lg-3 col-md-6">
            <div class="count-box">
              <i class="bi bi-journal-richtext" style="color: #ee6c20;"></i>
              <div>
                <span data-purecounter-start="0" data-purecounter-end="521" data-purecounter-duration="1" class="purecounter"></span>
                <p>Projects</p>
              </div>
            </div>
          </div>

          <div class="col-lg-3 col-md-6">
            <div class="count-box">
              <i class="bi bi-headset" style="color: #15be56;"></i>
              <div>
                <span data-purecounter-start="0" data-purecounter-end="1463" data-purecounter-duration="1" class="purecounter"></span>
                <p>Hours Of Support</p>
              </div>
            </div>
          </div>

          <div class="col-lg-3 col-md-6">
            <div class="count-box">
              <i class="bi bi-people" style="color: #bb0852;"></i>
              <div>
                <span data-purecounter-start="0" data-purecounter-end="5" data-purecounter-duration="1" class="purecounter"></span>
                <p>Hard Workers</p>
              </div>
            </div>
          </div>

        </div>

      </div>
    </section><!-- End Counts Section -->

    <!-- ======= Features Section ======= -->
    <section id="features" class="features">

      <div class="container" data-aos="fade-up">

        <header class="section-header">
          <h2>Features</h2>
          <p>Fully Responsive Eye-Catching Solutions</p>
        </header>

        <div class="row">

          <div class="col-lg-6">
            <img src="{{ asset('/assets/img/features.png') }}" class="img-fluid" alt="">
          </div>

          <div class="col-lg-6 mt-5 mt-lg-0 d-flex">
            <div class="row align-self-center gy-4">

              <div class="col-md-6" data-aos="zoom-out" data-aos-delay="200">
                <div class="feature-box d-flex align-items-center">
                  <i class="bi bi-check"></i>
                  <h3>WordPress Divi Theme</h3>
                </div>
              </div>

              <div class="col-md-6" data-aos="zoom-out" data-aos-delay="300">
                <div class="feature-box d-flex align-items-center">
                  <i class="bi bi-check"></i>
                  <h3>Laravel Framework 8.x</h3>
                </div>
              </div>

              <div class="col-md-6" data-aos="zoom-out" data-aos-delay="400">
                <div class="feature-box d-flex align-items-center">
                  <i class="bi bi-check"></i>
                  <h3>Bootstrap / HTML5 / CSS3</h3>
                </div>
              </div>

              <div class="col-md-6" data-aos="zoom-out" data-aos-delay="500">
                <div class="feature-box d-flex align-items-center">
                  <i class="bi bi-check"></i>
                  <h3>PHP 7.x</h3>
                </div>
              </div>

              <div class="col-md-6" data-aos="zoom-out" data-aos-delay="600">
                <div class="feature-box d-flex align-items-center">
                  <i class="bi bi-check"></i>
                  <h3>MySQL / PostgreSQL</h3>
                </div>
              </div>

              <div class="col-md-6" data-aos="zoom-out" data-aos-delay="700">
                <div class="feature-box d-flex align-items-center">
                  <i class="bi bi-check"></i>
                  <h3>JavaScript</h3>
                </div>
              </div>

            </div>
          </div>

        </div> <!-- / row -->

        <!-- Feature Tabs -->
        <div class="row feture-tabs" data-aos="fade-up">
          <div class="col-lg-6">
            <h3>Business Software & Web Application </h3>

            <!-- Tabs -->
            <ul class="nav nav-pills mb-3">
              <li>
                <a class="nav-link active" data-bs-toggle="pill" href="#tab1">Business Software</a>
              </li>
              <li>
                <a class="nav-link" data-bs-toggle="pill" href="#tab2">Web Application</a>
              </li>
              <li>
                <a class="nav-link" data-bs-toggle="pill" href="#tab3">Available Apps </a>
              </li>
            </ul><!-- End Tabs -->

            <!-- Tab Content -->
            <div class="tab-content">

              <div class="tab-pane fade show active" id="tab1">
                <p> There are many different types of business software. 
                  Deciding which one is best for your business is not always
                   easy and will often depend on:</p>
                <div class="d-flex align-items-center mb-2">
                  <i class="bi bi-check2"></i>
                  <h4>Common types of business software</h4>
                </div>
                <p>Common types of software used in business are:</p>
                <div class="d-flex align-items-center mb-2">

                  Accounts Software<br>
                  Billing Software<br>
                  Payroll Software<br>
                  Database Software<br>
                  Business Software<br>
                  Asset Management Software<br>
                  Desktop Publishing Programs<br>
              
                </div>
                <div class="d-flex align-items-center mb-2">
                  <i class="bi bi-check2"></i>
                  <h4>These usually come as bespoke products or off-the-shelf software.</h4>
                </div>
              </div><!-- End Tab 1 Content -->

              <div class="tab-pane fade show" id="tab2">
                <p>
                  Business web application systems for web and mobile will set your company apart from competition and increase business productivity. 
                 </p>
                <p>
                  A web application or "web app" is a software program that runs on a web server. 
                  Unlike traditional desktop applications, which are launched by your operating system, 
                  web apps must be accessed through a web browser...</p> 
                  <p>
                  Developers do not need to distribute software updates to users when the web app is updated.             

                </p>

              </div><!-- End Tab 2 Content -->

              <div class="tab-pane fade show" id="tab3">
                <p>
                  Business web application systems for web and mobile will set your company apart from competition and increase business productivity. 
                 </p>

                <p>Online Stores</p> 
                <p>Billing Application</p> 
                <p>Real Estate Application</p> 
                <p>Financial Institutions & Organizations</p> 
                <p>Schools and Higher Education</p> 
                <p>Corporations and Creative Agencies</p> 
            
              
              
              </div><!-- End Tab 3 Content -->

            </div>

          </div>

          <div class="col-lg-6">
            <img src="{{ asset('/assets/img/features-2.png') }}" class="img-fluid" alt="">
          </div>

        </div><!-- End Feature Tabs -->

        <!-- Feature Icons -->
        <div class="row feature-icons" data-aos="fade-up">
          <h3>Responsive Design</h3>

          <div class="row">

            <div class="col-xl-4 text-center" data-aos="fade-right" data-aos-delay="100">
              <img src="{{ asset('/assets/img/features-3.png') }}" class="img-fluid p-4" alt="">
            </div>

            <div class="col-xl-8 d-flex content">
              <div class="row align-self-center gy-4">


                <div class="col-md-6 icon-box" data-aos="fade-up" data-aos-delay="100">
                  <i class="ri-customer-service-2-line"></i>
                  <div>
                    <h4>Customer-centric solutions</h4>
                    <p>Visual hierarchy and progressive 
                      disclosure and navigation drawers to give users needed items first.</p>
                  </div>
                </div>

                <div class="col-md-6 icon-box" data-aos="fade-up" data-aos-delay="200">
                  <i class="ri-device-line"></i>
                  <div>
                    <h4>“Mobile first” approach</h4>
                    <p>Phone-sized designs to suit larger screens.</p>
                  </div>
                </div>

                <div class="col-md-6 icon-box" data-aos="fade-up">
                  <i class="ri-user-follow-line"></i>
                  <div>
                    <h4>Client-focused</h4>
                    <p>Consequuntur sunt aut quasi enim aliquam quae harum pariatur laboris nisi ut aliquip</p>
                  </div>
                </div>

                <div class="col-md-6 icon-box" data-aos="fade-up" data-aos-delay="300">
                  <i class="ri-collage-line"></i>
                  <div>
                    <h4>Fluid grids and images</h4>
                    <p>Images in native dimension, or croped for maximize impact</p>
                  </div>
                </div>

                <div class="col-md-6 icon-box" data-aos="fade-up" data-aos-delay="400">
                  <i class="ri-text-wrap"></i>
                  <div>
                    <h4>Users’ contexts</h4>
                    <p>Prioritize and hide content to suit users’ contexts</p>
                  </div>
                </div>

                <div class="col-md-6 icon-box" data-aos="fade-up" data-aos-delay="500">
                  <i class="ri-pencil-ruler-line"></i>
                  <div>
                    <h4>Design patterns</h4>
                    <p>Ease of use for users, quicken their familiarity</p>
                  </div>
                </div>

              </div>
            </div>

          </div>

        </div><!-- End Feature Icons -->

      </div>

    </section><!-- End Features Section -->

    <!-- ======= Services Section ======= -->
    <section id="services" class="services">

      <div class="container" data-aos="fade-up">

        <header class="section-header">
          <h2>Services</h2>
          <p>Website Design and Development</p>
        </header>

        <div class="row gy-4">

          <div class="col-lg-4 col-md-6" data-aos="fade-up" data-aos-delay="200">
            <div class="service-box blue">
              <i class="ri-layout-masonry-line icon"></i>
              <h3>Responsive Design</h3>
              <p>

                Client-focused, customer-centric solutions.</p>
              <a href="#" class="read-more"><span>Read More</span> <i class="bi bi-arrow-right"></i></a>
            </div>
          </div>

          <div class="col-lg-4 col-md-6" data-aos="fade-up" data-aos-delay="300">
            <div class="service-box orange">
              <i class="ri-coins-line icon"></i>
              <h3>E-Commerce</h3>
              <p>Start selling right away!</p>
              <a href="#" class="read-more"><span>Read More</span> <i class="bi bi-arrow-right"></i></a>
            </div>
          </div>

          <div class="col-lg-4 col-md-6" data-aos="fade-up" data-aos-delay="400">
            <div class="service-box green">
              <i class="ri-user-follow-line icon"></i>
              <h3>User Experience</h3>
              <p>Your customer will love it!</p>
              <a href="#" class="read-more"><span>Read More</span> <i class="bi bi-arrow-right"></i></a>
            </div>
          </div>

          <div class="col-lg-4 col-md-6" data-aos="fade-up" data-aos-delay="500">
            <div class="service-box red">
              <i class="ri-rocket-line icon"></i>
              <h3>Fast Page Loading</h3>
              <p>We don’t make our customers wait.</p>ri-run-line
              <a href="#" class="read-more"><span>Read More</span> <i class="bi bi-arrow-right"></i></a>
            </div>
          </div>

          <div class="col-lg-4 col-md-6" data-aos="fade-up" data-aos-delay="600">
            <div class="service-box purple">
              <i class="ri-wheelchair-line icon"></i>
              <h3>Aim for accessibility.</h3>
              <p>Contrast and background effectiveness, all text responsive</p>
              <a href="#" class="read-more"><span>Read More</span> <i class="bi bi-arrow-right"></i></a>
            </div>
          </div>

          <div class="col-lg-4 col-md-6" data-aos="fade-up" data-aos-delay="700">
            <div class="service-box pink">
              <i class="ri-device-line icon"></i> 
              <h3>Mobile solutions</h3>
              <p>Added value to mobile solutions for society, 
                industries, organizations, and administrative departments.</p>
              <a href="#" class="read-more"><span>Read More</span> <i class="bi bi-arrow-right"></i></a>
            </div>
          </div>

        </div>

      </div>

    </section><!-- End Services Section -->

    <!-- ======= Pricing Section ======= -->
    <section id="pricing" class="pricing">

      <div class="container" data-aos="fade-up">

        <header class="section-header">
          <h2>Pricing</h2>
          <p>Check our Pricing</p>
        </header>

        <div class="row gy-4" data-aos="fade-left">

          <div class="col-lg-3 col-md-6" data-aos="zoom-in" data-aos-delay="100">
            <div class="box">
              <h3 style="color: #07d5c0;">Free Quote</h3>
              <div class="price"><sup>$</sup>0<span> / mo</span></div>
              <img src="{{ asset('/assets/img/pricing-free.png') }}" class="img-fluid" alt="">
              <ul>
                <li>Web site app quote</li>
                <li>Revisions</li>
                <li class="na">Installation</li>
                <li class="na">Database design</li>
                <li class="na">Web Design</li>
              </ul>
              <a href="mailto:tn76contact@gmail.com?subject=Contacting throu tn76.com (Free Quote)&body=I need you to ..." class="btn-buy">Contact Us</a>
            </div>
          </div>

          <div class="col-lg-3 col-md-6" data-aos="zoom-in" data-aos-delay="200">
            <div class="box">
              <h3 style="color: #65c600;">Software Installation</h3>
              <div class="price">From <sup>$</sup>1<span></span></div>
              <img src="{{ asset('/assets/img/pricing-starter.png') }}" class="img-fluid" alt="">
              <ul>
                <li>Software Upload</li>
                <li>Installation</li>
                <li>Database setup + migr</li>
                <li class="na">Upgrades</li>
                <li class="na">Bug fixes</li>
              </ul>
              <a href="mailto:tn76contact@gmail.com?subject=Contacting throu tn76.com (Installation)&body=I need you to ..." class="btn-buy">Contact Us</a>
            </div>
          </div>

          <div class="col-lg-3 col-md-6" data-aos="zoom-in" data-aos-delay="300">
            <div class="box">
              <span class="featured">Featured</span>
              <h3 style="color: #ff901c;">Website/ app Creation</h3>
              <div class="price">From <sup>$</sup>10<span> +</span></div>
              <img src="{{ asset('/assets/img/pricing-business.png') }}" class="img-fluid" alt="">
              <ul>
                <li>Home page</li>
                <li>3 pages</li>
                <li>Contact form</li>
                <li>10 revisions</li>
                <li>Installation</li>
              </ul>
              <a href="mailto:tn76contact@gmail.com?subject=Contacting throu tn76.com (Featured)&body=I need you to ..." class="btn-buy">Contact Us</a>
            </div>
          </div>

          <div class="col-lg-3 col-md-6" data-aos="zoom-in" data-aos-delay="400">
            <div class="box">
              <h3 style="color: #ff0071;">Full Stack service</h3>
              <div class="price"><sup>$</sup>70<span> / h</span></div>
              <img src="{{ asset('/assets/img/pricing-ultimate.png') }}" class="img-fluid" alt="">
              <ul>
                <li>Website/ app Creation</li>
                <li>Hosting/Cloud/ installation</li>
                <li>Unlimited ugrades</li>
                <li>Unlimited upadtes</li>
                <li>Unlimited revisions</li>
              </ul>
              <a href="mailto:tn76contact@gmail.com?subject=Contacting throu tn76.com (Full Stack service)&body=I need you to ..." class="btn-buy">Contact Us</a>
            </div>
          </div>

        </div>

      </div>

    </section><!-- End Pricing Section -->

    <!-- ======= F.A.Q Section ======= -->
    <section id="faq" class="faq">

      <div class="container" data-aos="fade-up">

        <header class="section-header">
          <h2>F.A.Q</h2>
          <p>Frequently Asked Questions</p>
        </header>

        <div class="row">
          <div class="col-lg-6">
            <!-- F.A.Q List 1-->
            <div class="accordion accordion-flush" id="faqlist1">

              <div class="accordion-item">
                <h2 class="accordion-header">
                  <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#faq-content-2">
                    How much does a new website cost? 
                  </button>
                </h2>
                <div id="faq-content-2" class="accordion-collapse collapse" data-bs-parent="#faqlist1">
                  <div class="accordion-body">
                    There’s not one single answer to this question. 
                    A website design is quoted based on the needs of each 
                    individual project. Every website is unique and requires
                     different components; we design and develop custom websites 
                     specifically for your small business. We’ll ask a lot of questions, 
                     assess your needs, and give you a quote based on that assessment.
                      Most of our sites run in the $3,000 – $5,000 range, but can be more or 
                      less depending on needs.
                  </div>
                </div>
              </div>




              <div class="accordion-item">
                <h2 class="accordion-header">
                  <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#faq-content-1">
                    How long will it take to get a new website?
                  </button>
                </h2>
                <div id="faq-content-1" class="accordion-collapse collapse" data-bs-parent="#faqlist1">
                  <div class="accordion-body">
                    On average, we shoot for a six to eight week turnaround, 
                    but the pace of any project is set by each client. 
                    How much input you can provide during the initial stages, 
                    your availability with feedback, how soon the content is ready 
                    – all this affects the speed of completion. The functionality 
                    needs may also play a role – more complex sites will take more time to develop.
                  </div>
                </div>
              </div>

              <div class="accordion-item">
                <h2 class="accordion-header">
                  <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#faq-content-3">
                    Do you only create WordPress websites? 
                  </button>
                </h2>
                <div id="faq-content-3" class="accordion-collapse collapse" data-bs-parent="#faqlist1">
                  <div class="accordion-body">
                    Yes. WordPress is a great content management system that now powers 
                    more than 30% of the web. Its huge market share is due to 
                    it’s flexibility and ease of use. Don’t worry, we’ll teach you 
                    how to manage it and help you along the way.
                  </div>
                </div>
              </div>

            </div>
          </div>

          <div class="col-lg-6">

            <!-- F.A.Q List 2-->
            <div class="accordion accordion-flush" id="faqlist2">

              <div class="accordion-item">
                <h2 class="accordion-header">
                  <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#faq2-content-1">
                    What if I need help on my site down the road? 
                  </button>
                </h2>
                <div id="faq2-content-1" class="accordion-collapse collapse" data-bs-parent="#faqlist2">
                  <div class="accordion-body">
                    We are only an email away! We’re here to help you as much or as little as you need, and we won’t disappear once the site is launched. We’ve been doing this for many years, so we aren’t going anywhere anytime soon.
                  </div>
                </div>
              </div>

              <div class="accordion-item">
                <h2 class="accordion-header">
                  <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#faq2-content-2">
                    When do I pay?  
                  </button>
                </h2>
                <div id="faq2-content-2" class="accordion-collapse collapse" data-bs-parent="#faqlist2">
                  <div class="accordion-body">
                    For most projects, equal payments are made at the start, 
                    midway, and the end of the project, but we can work with you 
                    to set up a schedule that meets your needs. We understand that 
                    this is a big investment and want to help you 
                    budget for the expense in whatever way possible.<br>
                    We accept checks, PayPal, and all major credit cards.        
                  </div>
                </div>
              </div>

              <div class="accordion-item">
                <h2 class="accordion-header">
                  <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#faq2-content-3">
                    Will my website be mobile-friendly? 
                  </button>
                </h2>
                <div id="faq2-content-3" class="accordion-collapse collapse" data-bs-parent="#faqlist2">
                  <div class="accordion-body">
                  Absolutely! Having a mobile-friendly website is more important than ever! 
                  We work hard to ensure your website looks great on a variety of devices.
                  </div>
                </div>
              </div>

            </div>
          </div>

        </div>

      </div>

    </section><!-- End F.A.Q Section -->

    <!-- ======= Portfolio Section ======= -->
    <section id="portfolio" class="portfolio">

      <div class="container" data-aos="fade-up">

        <header class="section-header">
          <h2>Portfolio</h2>
          <p>Check our latest work</p>
        </header>

        <div class="row" data-aos="fade-up" data-aos-delay="100">
          <div class="col-lg-12 d-flex justify-content-center">
            <ul id="portfolio-flters">
              <li data-filter="*" class="filter-active">All</li>
              <li data-filter=".filter-app">App</li>
              <li data-filter=".filter-card">Card</li>
              <li data-filter=".filter-web">Web</li>
            </ul>
          </div>
        </div>

        <div class="row gy-4 portfolio-container" data-aos="fade-up" data-aos-delay="200">

          <div class="col-lg-4 col-md-6 portfolio-item filter-web">
            <div class="portfolio-wrap">
              <img src="{{ asset('/assets/img/Portfolio/img/Medilab-bootstrap-website-template.png') }}" class="img-fluid" alt="">
              <div class="portfolio-info">
                <h4>Medilab</h4>
                <p>Web</p>
                <div class="portfolio-links">
                  <a href="{{ asset('/assets/img/Portfolio/img/Medilab-bootstrap-website-template.png') }}" data-gallery="portfolioGallery" class="portfokio-lightbox" title="Medilab"><i class="bi bi-plus"></i></a>
                  <a href="Portfolio/medilab.html" title="Demo"  target="_Blank"><i class="bi bi-link"></i></a>
                </div>
              </div>
            </div>
          </div>

          <div class="col-lg-4 col-md-6 portfolio-item filter-web">
            <div class="portfolio-wrap">
              <img src="{{ asset('/assets/img/Portfolio/img/OnePage-bootstrap-website-template.png') }}" class="img-fluid" alt="">
              <div class="portfolio-info">
                <h4>OnePage</h4>
                <p>Web</p>
                <div class="portfolio-links">
                  <a href="{{ asset('/assets/img/portfolio/img/OnePage-bootstrap-website-template.png') }}" data-gallery="portfolioGallery" class="portfokio-lightbox" title="One Page"><i class="bi bi-plus"></i></a>
                  <a href="https://bootstrapmade.com/demo/templates/OnePage/" title="Demo"  target="_Blank"><i class="bi bi-link"></i></a>
                </div>
              </div>
            </div>
          </div>

          <div class="col-lg-4 col-md-6 portfolio-item filter-app">
            <div class="portfolio-wrap">
              <img src="{{ asset('/assets/img/portfolio/img/portfolio-3.jpg') }}" class="img-fluid" alt="">
              <div class="portfolio-info">
                <h4>App 2</h4>
                <p>App</p>
                <div class="portfolio-links">
                  <a href="{{ asset('/assets/img/portfolio/img/portfolio-3.jpg') }}" data-gallery="portfolioGallery" class="portfokio-lightbox" title="App 2"><i class="bi bi-plus"></i></a>
                  <a href="portfolio-details.html" title="More Details"><i class="bi bi-link"></i></a>
                </div>
              </div>
            </div>
          </div>

          <div class="col-lg-4 col-md-6 portfolio-item filter-card">
            <div class="portfolio-wrap">
              <img src="{{ asset('/assets/img/portfolio/img/portfolio-4.jpg') }}" class="img-fluid" alt="">
              <div class="portfolio-info">
                <h4>Card 2</h4>
                <p>Card</p>
                <div class="portfolio-links">
                  <a href="{{ asset('/assets/img/portfolio/img/portfolio-4.jpg') }}" data-gallery="portfolioGallery" class="portfokio-lightbox" title="Card 2"><i class="bi bi-plus"></i></a>
                  <a href="portfolio-details.html" title="More Details"><i class="bi bi-link"></i></a>
                </div>
              </div>
            </div>
          </div>

          <div class="col-lg-4 col-md-6 portfolio-item filter-web">
            <div class="portfolio-wrap">
              <img src="{{ asset('/assets/img/portfolio/img/Company-bootstrap-website-template.png') }}" class="img-fluid" alt="">
              <div class="portfolio-info">
                <h4>Company</h4>
                <p>Web</p>
                <div class="portfolio-links">
                  <a href="{{ asset('/assets/img/Portfolio/Company-bootstrap-website-template.png') }}" data-gallery="portfolioGallery" class="portfokio-lightbox" title="Company"><i class="bi bi-plus"></i></a>
                  <a href="https://bootstrapmade.com/demo/templates/Company/index.html" title="Demo" target="_Blank"><i class="bi bi-link"></i></a>
                </div>
              </div>
            </div>
          </div>

          <div class="col-lg-4 col-md-6 portfolio-item filter-app">
            <div class="portfolio-wrap">
              <img src="{{ asset('/assets/img/portfolio/img/portfolio-6.jpg') }}" class="img-fluid" alt="">
              <div class="portfolio-info">
                <h4>App 3</h4>
                <p>App</p>
                <div class="portfolio-links">
                  <a href="{{ asset('/assets/img/portfolio/img/portfolio-6.jpg') }}" data-gallery="portfolioGallery" class="portfokio-lightbox" title="App 3"><i class="bi bi-plus"></i></a>
                  <a href="portfolio-details.html" title="More Details"><i class="bi bi-link"></i></a>
                </div>
              </div>
            </div>
          </div>

          <div class="col-lg-4 col-md-6 portfolio-item filter-card">
            <div class="portfolio-wrap">
              <img src="{{ asset('/assets/img/portfolio/img/portfolio-7.jpg') }}" class="img-fluid" alt="">
              <div class="portfolio-info">
                <h4>Card 1</h4>
                <p>Card</p>
                <div class="portfolio-links">
                  <a href="{{ asset('/assets/img/portfolio/img/portfolio-7.jpg') }}" data-gallery="portfolioGallery" class="portfokio-lightbox" title="Card 1"><i class="bi bi-plus"></i></a>
                  <a href="portfolio-details.html" title="More Details"><i class="bi bi-link"></i></a>
                </div>
              </div>
            </div>
          </div>

          <div class="col-lg-4 col-md-6 portfolio-item filter-card">
            <div class="portfolio-wrap">
              <img src="{{ asset('/assets/img/portfolio/img/portfolio-8.jpg') }}" class="img-fluid" alt="">
              <div class="portfolio-info">
                <h4>Card 3</h4>
                <p>Card</p>
                <div class="portfolio-links">
                  <a href="{{ asset('/assets/img/portfolio/img/portfolio-8.jpg') }}" data-gallery="portfolioGallery" class="portfokio-lightbox" title="Card 3"><i class="bi bi-plus"></i></a>
                  <a href="portfolio-details.html" title="More Details"><i class="bi bi-link"></i></a>
                </div>
              </div>
            </div>
          </div>

          <div class="col-lg-4 col-md-6 portfolio-item filter-web">
            <div class="portfolio-wrap">
              <img src="{{ asset('/assets/img/portfolio/img/portfolio-9.jpg') }}" class="img-fluid" alt="">
              <div class="portfolio-info">
                <h4>Web 3</h4>
                <p>Web</p>
                <div class="portfolio-links">
                  <a href="{{ asset('/assets/img/portfolio/img/portfolio-9.jpg') }}" data-gallery="portfolioGallery" class="portfokio-lightbox" title="Web 3"><i class="bi bi-plus"></i></a>
                  <a href="portfolio-details.html" title="More Details"><i class="bi bi-link"></i></a>
                </div>
              </div>
            </div>
          </div>

        </div>

      </div>

    </section><!-- End Portfolio Section -->

    <!-- ======= Testimonials Section ======= -->
    <section id="testimonials" class="testimonials">

      <div class="container" data-aos="fade-up">

        <header class="section-header">
          <h2>Testimonials</h2>
          <p>What our clients are saying about us</p>
        </header>

        <div class="testimonials-slider swiper-container" data-aos="fade-up" data-aos-delay="200">
          <div class="swiper-wrapper">

            <div class="swiper-slide">
              <div class="testimonial-item">
                <div class="stars">
                  <i class="bi bi-star-fill"></i><i class="bi bi-star-fill"></i><i class="bi bi-star-fill"></i><i class="bi bi-star-fill"></i><i class="bi bi-star-fill"></i>
                </div>
                <p>
                  @tn76com my best experience ever.
                </p>
                <div class="profile mt-auto">
                  <img src="{{ asset('/assets/img/testimonials/testimonials-1.jpg') }}" class="testimonial-img" alt="">
                  <h3>Saul Goodman</h3>
                  <h4>Ceo &amp; Founder</h4>
                </div>
              </div>
            </div><!-- End testimonial item -->

            <div class="swiper-slide">
              <div class="testimonial-item">
                <div class="stars">
                  <i class="bi bi-star-fill"></i><i class="bi bi-star-fill"></i><i class="bi bi-star-fill"></i><i class="bi bi-star-fill"></i><i class="bi bi-star-fill"></i>
                </div>
                <p>
                  My website built using Wordpress Divi theme is SO GOOD 100/10.
                </p>
                <div class="profile mt-auto">
                  <img src="{{ asset('/assets/img/testimonials/testimonials-2.jpg') }}" class="testimonial-img" alt="">
                  <h3>Sara Wilsson</h3>
                  <h4>3F Founder</h4>
                </div>
              </div>
            </div><!-- End testimonial item -->

            <div class="swiper-slide">
              <div class="testimonial-item">
                <div class="stars">
                  <i class="bi bi-star-fill"></i><i class="bi bi-star-fill"></i><i class="bi bi-star-fill"></i><i class="bi bi-star-fill"></i><i class="bi bi-star-fill"></i>
                </div>
                <p>
                  We're running all tiketing trough their awesome platform for our event #fwssummit                
                </p>
                <div class="profile mt-auto">
                  <img src="{{ asset('/assets/img/testimonials/testimonials-3.jpg') }}" class="testimonial-img" alt="">
                  <h3>Jena Karlis</h3>
                  <h4>Event manager</h4>
                </div>
              </div>
            </div><!-- End testimonial item -->

            <div class="swiper-slide">
              <div class="testimonial-item">
                <div class="stars">
                  <i class="bi bi-star-fill"></i><i class="bi bi-star-fill"></i><i class="bi bi-star-fill"></i><i class="bi bi-star-fill"></i><i class="bi bi-star-fill"></i>
                </div>
                <p>
                  @tn76com I love my @pizzagotn :) Easy platform to use, fantastic staff and nothing but greay results!
                </p>
                <div class="profile mt-auto">
                  <img src="{{ asset('/assets/img/testimonials/testimonials-4.jpg') }}" class="testimonial-img" alt="">
                  <h3>Matt Brandon</h3>
                  <h4>Restaurant owner </h4>
                </div>
              </div>
            </div><!-- End testimonial item -->

            <div class="swiper-slide">
              <div class="testimonial-item">
                <div class="stars">
                  <i class="bi bi-star-fill"></i><i class="bi bi-star-fill"></i><i class="bi bi-star-fill"></i><i class="bi bi-star-fill"></i><i class="bi bi-star-fill"></i>
                </div>
                <p>
                  @tn76com So glad to work with you.
                </p>
                <div class="profile mt-auto">
                  <img src="{{ asset('/assets/img/testimonials/testimonials-5.jpg') }}" class="testimonial-img" alt="">
                  <h3>John Larson</h3>
                  <h4>Entrepreneur</h4>
                </div>
              </div>
            </div><!-- End testimonial item -->

          </div>
          <div class="swiper-pagination"></div>
        </div>

      </div>

    </section><!-- End Testimonials Section -->

    <!-- ======= Team Section ======= -->
    <section id="team" class="team">

      <div class="container" data-aos="fade-up">

        <header class="section-header">
          <h2>Team</h2>
          <p>Our hard working team</p>
        </header>

        <div class="row gy-4">

          <div class="col-lg-3 col-md-6 d-flex align-items-stretch" data-aos="fade-up" data-aos-delay="100">
            <div class="member">
              <div class="member-img">
                <img src="{{ asset('/assets/img/team/team-1.jpg') }}" class="img-fluid" alt="">
                <div class="social">
                  <a href=""><i class="bi bi-twitter"></i></a>
                  <a href=""><i class="bi bi-facebook"></i></a>
                  <a href=""><i class="bi bi-instagram"></i></a>
                  <a href=""><i class="bi bi-linkedin"></i></a>
                </div>
              </div>
              <div class="member-info">
                <h4>Walter White</h4>
                <span>Chief Executive Officer</span>
                <p>Velit aut quia fugit et et. Dolorum ea voluptate vel tempore tenetur ipsa quae aut. Ipsum exercitationem iure minima enim corporis et voluptate.</p>
              </div>
            </div>
          </div>

          <div class="col-lg-3 col-md-6 d-flex align-items-stretch" data-aos="fade-up" data-aos-delay="200">
            <div class="member">
              <div class="member-img">
                <img src="{{ asset('/assets/img/team/team-2.jpg') }}" class="img-fluid" alt="">
                <div class="social">
                  <a href=""><i class="bi bi-twitter"></i></a>
                  <a href=""><i class="bi bi-facebook"></i></a>
                  <a href=""><i class="bi bi-instagram"></i></a>
                  <a href=""><i class="bi bi-linkedin"></i></a>
                </div>
              </div>
              <div class="member-info">
                <h4>Sarah Jhonson</h4>
                <span>Product Manager</span>
                <p>Quo esse repellendus quia id. Est eum et accusantium pariatur fugit nihil minima suscipit corporis. Voluptate sed quas reiciendis animi neque sapiente.</p>
              </div>
            </div>
          </div>

          <div class="col-lg-3 col-md-6 d-flex align-items-stretch" data-aos="fade-up" data-aos-delay="300">
            <div class="member">
              <div class="member-img">
                <img src="{{ asset('/assets/img/team/team-3.jpg') }}" class="img-fluid" alt="">
                <div class="social">
                  <a href=""><i class="bi bi-twitter"></i></a>
                  <a href=""><i class="bi bi-facebook"></i></a>
                  <a href=""><i class="bi bi-instagram"></i></a>
                  <a href=""><i class="bi bi-linkedin"></i></a>
                </div>
              </div>
              <div class="member-info">
                <h4>William Anderson</h4>
                <span>CTO</span>
                <p>Vero omnis enim consequatur. Voluptas consectetur unde qui molestiae deserunt. Voluptates enim aut architecto porro aspernatur molestiae modi.</p>
              </div>
            </div>
          </div>

          <div class="col-lg-3 col-md-6 d-flex align-items-stretch" data-aos="fade-up" data-aos-delay="400">
            <div class="member">
              <div class="member-img">
                <img src="{{ asset('/assets/img/team/team-4.jpg') }}" class="img-fluid" alt="">
                <div class="social">
                  <a href=""><i class="bi bi-twitter"></i></a>
                  <a href=""><i class="bi bi-facebook"></i></a>
                  <a href=""><i class="bi bi-instagram"></i></a>
                  <a href=""><i class="bi bi-linkedin"></i></a>
                </div>
              </div>
              <div class="member-info">
                <h4>Amanda Jepson</h4>
                <span>Accountant</span>
                <p>Rerum voluptate non adipisci animi distinctio et deserunt amet voluptas. Quia aut aliquid doloremque ut possimus ipsum officia.</p>
              </div>
            </div>
          </div>

        </div>

      </div>

    </section><!-- End Team Section -->

    <!-- ======= Clients Section ======= -->
    <section id="clients" class="clients">

      <div class="container" data-aos="fade-up">

        <header class="section-header">
          <h2>Our Clients</h2>
          <p>Committed to reshaping their organizations and driving lasting value</p>
        </header>

        <div class="clients-slider swiper-container">
          <div class="swiper-wrapper align-items-center">
            <div class="swiper-slide"><img src="{{ asset('/assets/img/clients/client-1.png') }}" class="img-fluid" alt=""></div>
            <div class="swiper-slide"><img src="{{ asset('/assets/img/clients/client-2.png') }}" class="img-fluid" alt=""></div>
            <div class="swiper-slide"><img src="{{ asset('/assets/img/clients/client-3.png') }}" class="img-fluid" alt=""></div>
            <div class="swiper-slide"><img src="{{ asset('/assets/img/clients/client-4.png') }}" class="img-fluid" alt=""></div>
            <div class="swiper-slide"><img src="{{ asset('/assets/img/clients/client-5.png') }}" class="img-fluid" alt=""></div>
            <div class="swiper-slide"><img src="{{ asset('/assets/img/clients/client-6.png') }}" class="img-fluid" alt=""></div>
            <div class="swiper-slide"><img src="{{ asset('/assets/img/clients/client-7.png') }}" class="img-fluid" alt=""></div>
            <div class="swiper-slide"><img src="{{ asset('/assets/img/clients/client-8.png') }}" class="img-fluid" alt=""></div>
          </div>
          <div class="swiper-pagination"></div>
        </div>
      </div>

    </section><!-- End Clients Section -->

    <!-- ======= Recent Blog Posts Section ======= -->
    <section id="recent-blog-posts" class="recent-blog-posts">

      <div class="container" data-aos="fade-up">

        <header class="section-header">
          <h2>Blog</h2>
          <p>Recent posts from our Blog</p>
        </header>

        <div class="row">

          
            
          <div class="col-lg-4">
            <div class="post-box">
              <div class="post-img"><img src="{{ asset('/assets/img/blog/blog-1.jpg') }}" class="img-fluid" alt=""></div>
              <span id="date1" name="date1" ></span>
              <h3 class="post-title">Eum ad dolor et. Autem aut fugiat debitis voluptatem consequuntur sit</h3>
              <a href="#" class="readmore stretched-link mt-auto"><span>Read More</span><i class="bi bi-arrow-right"></i></a>
            </div>
          </div>
          <div class="col-lg-4">
            <div class="post-box">
              <div class="post-img"><img src="{{ asset('/assets/img/blog/blog-2.jpg') }}" class="img-fluid" alt=""></div>
              <span id="date2" name="date2" ></span>
              <h3 class="post-title">Et repellendus molestiae qui est sed omnis voluptates magnam</h3>
              <a href="#" class="readmore stretched-link mt-auto"><span>Read More</span><i class="bi bi-arrow-right"></i></a>
            </div>
          </div>

          <div class="col-lg-4">
            <div class="post-box">
              <div class="post-img"><img src="{{ asset('/assets/img/blog/blog-3.jpg') }}" class="img-fluid" alt=""></div>
              <span id="date3" name="date3" ></span>
              <h3 class="post-title">Quia assumenda est et veritatis aut quae</h3>
              <a href="#" class="readmore stretched-link mt-auto"><span>Read More</span><i class="bi bi-arrow-right"></i></a>
            </div>
          </div>

        </div>
<script>
            const today = new Date()
            const date1 = new Date(today)
            const date2 = new Date(today)
            const date3 = new Date(today)
            date1.setDate(date1.getDate() - 1)
            date2.setDate(date2.getDate() - 1)
            date3.setDate(date3.getDate() - 2)
            document.getElementById("date1").innerHTML = date1.toDateString();
            document.getElementById("date2").innerHTML = date2.toDateString();
            document.getElementById("date3").innerHTML = date3.toDateString();
</script>

      </div>

    </section><!-- End Recent Blog Posts Section -->

    <!-- ======= Contact Section ======= -->
    <section id="contact" class="contact">

      <div class="container" data-aos="fade-up">

        <header class="section-header">
          <h2>Contact</h2>
          <p>Contact Us</p>
        </header>

        <div class="row gy-4">

          <div class="col-lg-12">

            <div class="row gy-4">
              <div class="col-md-6">
                <div class="info-box">
                  <i class="bi bi-geo-alt"></i>
                  <h3>Address</h3>
                  <p>
                    A201 Res Shopping <br>
                    New Tn, TN 532045<br>
                    Tunisia <br><br>
                  </p>
                </div>
              </div>
              <div class="col-md-6">
                <div class="info-box">
                  <i class="bi bi-chat-text"></i>
                  
                  <h3>Text Us</h3>
                  <p>
                    <a href="https://wa.me/+21697538607/?text=Hello, Please ">WhatsApp</a>
                  </p>
                  <p>                   
                    <a href="tel:+18312045509">+1 831 2045509</a>
                  </p>
                  
                </div>
              </div>
              <div class="col-md-6">
                <div class="info-box">
                  <i class="bi bi-envelope"></i>
                  <h3>Email Us</h3>
                  <p><a href="mailto:tn76contact@gmail.com?subject=Contacting throu tn76.com&body=I need you to ...">tn76contact@gmail.com</a>
				  </p>
                </div>
              </div>
              <div class="col-md-6">
                <div class="info-box">
                  <i class="bi bi-clock"></i>
                  <h3>Open Hours</h3>
                  <p>Monday - Friday<br>9:00AM - 05:00PM</p>
                </div>
              </div>
            </div>

          </div>



          </div>

        </div>

      </div>

    </section><!-- End Contact Section -->

  </main>
  <!-- End #main -->

  <!-- End reusable part -->

@endsection