@extends('layouts.main')

@section('title', 'Users Home')

@section('content')
 
<main id="main">
<hr>

    <!-- ======= user Section ======= -->
    <section id="user" class="user">

      <div class="container" data-aos="fade-up">

        <header class="section-header">
          <h2>Users</h2>
          <p>listing</p>
        </header>

        @if (session('successMsg'))
        <div class="alert alert-success" role="alert">
            {{ session('SuccessMsg') }}
        </div>



      </div>

    </section><!-- End user Section -->



</main><!-- End #main -->


@endsection