@extends('layouts.main')

@section('title', 'Home')

@section('content')

<main id="main">

      @if($errors->any())
      @foreach ($errors->all() as $error)
      <div class="alert alert-danger" role="alert">
        {{ $error }}
      </div>
      @endforeach
      @endif


    <!-- ======= user Section ======= -->
    <section id="user" class="user">

      <div class="container" data-aos="fade-up">





        <header class="section-header">
          <h2>Users</h2>
          <p>Create page</p>
        </header>


          <div class="col-lg-12">
            <form action="{{ Route('users') }}" method="post"> class="php-email-form"
            {{ csrf_field() }}  
            <div class="row gy-4">

                <div class="col-md-6">
                  <input type="text" name="first_name" class="form-control" placeholder="Name" >
                </div>

                <div class="col-md-6">
                  <input type="text" name="last_name" class="form-control" placeholder="Last Name" >
                </div>

                <div class="col-md-12 ">
                  <input type="email" class="form-control" name="email" placeholder="Email" >
                </div>

                <div class="col-md-12">
                  <input type="text" class="form-control" name="phone" placeholder="Phone" >
                </div>

                <div class="col-md-12 text-center">
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