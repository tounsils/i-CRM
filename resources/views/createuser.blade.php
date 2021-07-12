@extends('layouts.main')

@section('title', 'Home')

@section('content')

<main id="main">
<hr>

    <!-- ======= user Section ======= -->
    <section id="user" class="user">

      <div class="container" data-aos="fade-up">

        <header class="section-header">
          <h2>Users</h2>
          <p>Add new user</p>
        </header>


          <div class="col-lg-12">
            <form action="{{ storeuser }}" method="post" class="php-email-form">
              <div class="row gy-4">

              <div class="col-md-6">
                  <input type="text" name="first_name" class="form-control" placeholder="Name" required>
                </div>

                <div class="col-md-6">
                  <input type="text" name="last_name" class="form-control" placeholder="Last Name" required>
                </div>

                <div class="col-md-12 ">
                  <input type="email" class="form-control" name="email" placeholder="Email" required>
                </div>

                <div class="col-md-12">
                  <input type="text" class="form-control" name="phone" placeholder="Phone" required>
                </div>


                <div class="col-md-12 text-center">
                  <div class="loading">Loading</div>
                  <div class="error-message"></div>
                  <div class="sent-message">Your message has been sent. Thank you!</div>

                  <button type="submit">Add User</button>
                </div>
                
              </div>
            </form>

          </div>

        </div>

      </div>

    </section><!-- End user Section -->



</main><!-- End #main -->

@endsection