@extends('layouts.main')

@section('title', 'Company Home')

@section('content')
<main id="main">

    <!-- ======= company Section ======= -->
    <section id="company" class="company">

      <div class="container" data-aos="fade-up">

        <header class="section-header">
          <h2>Company</h2>
          <p>Listing</p>
        </header>



<main class="signup-form">
    <div class="cotainer">
        <div class="row justify-content-center">
            <div class="col-md-4">
                <div class="card">
                    <h3 class="card-header text-center">Register User</h3>
                    <div class="card-body">

                        <form action="{{ route('register.custom') }}" method="POST">
                        @csrf
                            <div class="form-group mb-3">
                                <input type="text" placeholder="first_name" id="first_name" class="form-control" name="first_name"
                                    required autofocus>
                                @if ($errors->has('first_name'))
                                <span class="text-danger">{{ $errors->first('first_name') }}</span>
                                @endif
                            </div>
                            <div class="form-group mb-3">
                                <input type="text" placeholder="last_name" id="last_name" class="form-control" name="last_name"
                                    required autofocus>
                                @if ($errors->has('last_name'))
                                <span class="text-danger">{{ $errors->first('last_name') }}</span>
                                @endif
                            </div>


                            <div class="form-group mb-3">
                                <input type="text" placeholder="Email" id="email" class="form-control"
                                    name="email" required autofocus>
                                @if ($errors->has('email'))
                                <span class="text-danger">{{ $errors->first('email') }}</span>
                                @endif
                            </div>
                            <div class="form-group mb-3">
                                <input type="text" placeholder="phone" id="phone" class="form-control"
                                    name="phone" required autofocus>
                                @if ($errors->has('phone'))
                                <span class="text-danger">{{ $errors->first('phone') }}</span>
                                @endif
                            </div>

                            <div class="form-group mb-3">
                                <input type="password" placeholder="Password" id="password" class="form-control"
                                    name="password" required>
                                @if ($errors->has('password'))
                                <span class="text-danger">{{ $errors->first('password') }}</span>
                                @endif
                            </div>

                            <div class="form-group mb-3">
                                <div class="checkbox">
                                    <label><input type="checkbox" name="remember"> Remember Me</label>
                                </div>
                            </div>

                            <div class="d-grid mx-auto">
                                <button type="submit" class="btn btn-dark btn-block">Sign up</button>
                            </div>
                        </form>

                    </div>
                </div>
            </div>
        </div>
    </div>
</main>

</div>

</section><!-- End company Section -->



</main>

@endsection