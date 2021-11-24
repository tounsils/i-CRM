@extends('layouts.main')

@section('title', 'Edit a Company ')

@section('content')

<main id="main">
<hr>

    <!-- ======= company Section ======= -->
    <section id="company" class="company">

      <div class="container" data-aos="fade-up">

        <header class="section-header">
          <h2>Company</h2>
          <p>Listing</p>
        </header>

        <div class="container">
        


<div class="row">
    <div class="col-sm-8 offset-sm-2">
        <h1 class="display-3">Update a company</h1>

        @if ($errors->any())
        <div class="alert alert-danger">
            <ul>
                @foreach ($errors->all() as $error)
                <li>{{ $error }}</li>
                @endforeach
            </ul>
        </div>
        <br /> 
        @endif
        <form method="post" action="{{ route('company.update', $company->id) }}">
            @method('PATCH') 
            @csrf
            <div class="form-group">

                <label for="company_name">Company Name:</label>
                <input type="text" class="form-control" name="company_name" value="{{ $company->company_name }}" />
            </div>

            <div class="form-group">
                <label for="email">Email:</label>
                <input type="text" class="form-control" name="email" value="{{ $company->email }}" />
            </div>

            <div class="form-group">
                <label for="phone">Phone:</label>
                <input type="text" class="form-control" name="phone" value="{{ $company->phone }}" />
            </div>
            <button type="submit" class="btn btn-primary">Update</button>
        </form>
    </div>
</div>


</div></div></section></main>

@endsection

@section('footer-scripts')
    @include('scripts.ajax')
@endsection